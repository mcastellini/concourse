module Build.Build exposing
    ( changeToBuild
    , documentTitle
    , getScrollBehavior
    , getUpdateMessage
    , handleCallback
    , handleDelivery
    , init
    , subscriptions
    , update
    , view
    )

import Build.Models exposing (BuildPageType(..), CurrentBuild, Model)
import Build.Output.Models exposing (OutputModel)
import Build.Output.Output
import Build.StepTree.Models as STModels
import Build.StepTree.StepTree as StepTree
import Build.Styles as Styles
import Concourse
import Concourse.BuildStatus
import Concourse.Pagination exposing (Paginated)
import DateFormat
import Dict exposing (Dict)
import EffectTransformer exposing (ET)
import Html exposing (Html)
import Html.Attributes
    exposing
        ( attribute
        , class
        , classList
        , href
        , id
        , style
        , title
        )
import Html.Events exposing (onBlur, onFocus, onMouseEnter, onMouseLeave)
import Html.Lazy
import Http
import Keyboard
import List.Extra
import Login.Login as Login
import Maybe.Extra
import Message.Callback exposing (Callback(..))
import Message.Effects as Effects exposing (Effect(..), ScrollDirection(..))
import Message.Message exposing (Hoverable(..), Message(..))
import Message.Subscription as Subscription exposing (Delivery(..), Interval(..), Subscription(..))
import Message.TopLevelMessage exposing (TopLevelMessage(..))
import RemoteData
import Routes
import StrictEvents exposing (onLeftClick, onMouseWheel)
import String
import Time
import UpdateMsg exposing (UpdateMsg)
import UserState exposing (UserState)
import Views.BuildDuration as BuildDuration
import Views.Icon as Icon
import Views.LoadingIndicator as LoadingIndicator
import Views.NotAuthorized as NotAuthorized
import Views.Spinner as Spinner
import Views.Styles
import Views.TopBar as TopBar


type alias Flags =
    { highlight : Routes.Highlight
    , pageType : BuildPageType
    }


type ScrollBehavior
    = ScrollWindow
    | NoScroll


init : Flags -> ( Model, List Effect )
init flags =
    changeToBuild
        flags
        ( { page = flags.pageType
          , now = Nothing
          , disableManualTrigger = False
          , history = []
          , nextPage = Nothing
          , currentBuild = RemoteData.NotAsked
          , browsingIndex = 0
          , autoScroll = True
          , previousKeyPress = Nothing
          , previousTriggerBuildByKey = False
          , showHelp = False
          , highlight = flags.highlight
          , hoveredElement = Nothing
          , hoveredCounter = 0
          , authorized = True
          , fetchingHistory = False
          , scrolledToCurrentBuild = False
          , shiftDown = False
          , isUserMenuExpanded = False
          }
        , [ GetCurrentTime ]
        )


subscriptions : Model -> List Subscription
subscriptions model =
    let
        buildEventsUrl =
            model.currentBuild
                |> RemoteData.toMaybe
                |> Maybe.andThen .output
                |> Maybe.andThen .eventStreamUrlPath
    in
    [ OnClockTick OneSecond
    , OnScrollToBottom
    , OnKeyDown
    , OnKeyUp
    , OnElementVisible
    ]
        ++ (case buildEventsUrl of
                Nothing ->
                    []

                Just url ->
                    [ Subscription.FromEventSource ( url, [ "end", "event" ] ) ]
           )


changeToBuild : Flags -> ET Model
changeToBuild { highlight, pageType } ( model, effects ) =
    if model.browsingIndex > 0 && pageType == model.page then
        ( { model | highlight = highlight }, effects )

    else
        let
            newIndex =
                model.browsingIndex + 1

            newBuild =
                RemoteData.map
                    (\cb -> { cb | prep = Nothing, output = Nothing })
                    model.currentBuild
        in
        ( { model
            | browsingIndex = newIndex
            , currentBuild = newBuild
            , autoScroll = True
            , page = pageType
            , highlight = highlight
          }
        , case pageType of
            OneOffBuildPage buildId ->
                effects
                    ++ [ CloseBuildEventStream
                       , FetchBuild 0 newIndex buildId
                       ]

            JobBuildPage jbi ->
                effects
                    ++ [ CloseBuildEventStream
                       , FetchJobBuild newIndex jbi
                       ]
        )


extractTitle : Model -> String
extractTitle model =
    case ( model.currentBuild |> RemoteData.toMaybe, currentJob model, model.page ) of
        ( Just build, Just { jobName }, _ ) ->
            jobName ++ " #" ++ build.build.name

        ( _, _, JobBuildPage { jobName, buildName } ) ->
            jobName ++ " #" ++ buildName

        ( _, _, OneOffBuildPage id ) ->
            "#" ++ String.fromInt id


currentJob : Model -> Maybe Concourse.JobIdentifier
currentJob =
    .currentBuild
        >> RemoteData.toMaybe
        >> Maybe.map .build
        >> Maybe.andThen .job


getUpdateMessage : Model -> UpdateMsg
getUpdateMessage model =
    case model.currentBuild of
        RemoteData.Failure _ ->
            UpdateMsg.NotFound

        _ ->
            UpdateMsg.AOK


handleCallback : Callback -> ET Model
handleCallback action ( model, effects ) =
    case action of
        BuildTriggered (Ok build) ->
            ( { model | history = build :: model.history }
            , effects
                ++ [ NavigateTo <| Routes.toString <| Routes.buildRoute build ]
            )

        BuildFetched (Ok ( browsingIndex, build )) ->
            handleBuildFetched browsingIndex build ( model, effects )

        BuildFetched (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( model, effects ++ [ RedirectToLogin ] )

                    else if status.code == 404 then
                        ( { model | currentBuild = RemoteData.Failure err }
                        , effects
                        )

                    else
                        ( model, effects )

                _ ->
                    ( model, effects )

        BuildAborted (Ok ()) ->
            ( model, effects )

        BuildPrepFetched (Ok ( browsingIndex, buildPrep )) ->
            handleBuildPrepFetched browsingIndex buildPrep ( model, effects )

        BuildPrepFetched (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( { model | authorized = False }, effects )

                    else
                        ( model, effects )

                _ ->
                    ( model, effects )

        PlanAndResourcesFetched buildId (Ok planAndResources) ->
            updateOutput
                (Build.Output.Output.planAndResourcesFetched
                    buildId
                    planAndResources
                )
                ( model
                , effects
                    ++ [ Effects.OpenBuildEventStream
                            { url =
                                "/api/v1/builds/"
                                    ++ String.fromInt buildId
                                    ++ "/events"
                            , eventTypes = [ "end", "event" ]
                            }
                       ]
                )

        PlanAndResourcesFetched buildId (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 404 then
                        let
                            url =
                                "/api/v1/builds/"
                                    ++ String.fromInt buildId
                                    ++ "/events"
                        in
                        updateOutput
                            (\m ->
                                ( { m | eventStreamUrlPath = Just url }
                                , []
                                , Build.Output.Output.OutNoop
                                )
                            )
                            ( model, effects )

                    else if status.code == 401 then
                        ( { model | authorized = False }, effects )

                    else
                        ( model, effects )

                _ ->
                    ( model, effects )

        BuildHistoryFetched (Ok history) ->
            handleHistoryFetched history ( model, effects )

        BuildHistoryFetched (Err _) ->
            -- https://github.com/concourse/concourse/issues/3201
            ( { model | fetchingHistory = False }, effects )

        BuildJobDetailsFetched (Ok job) ->
            ( { model | disableManualTrigger = job.disableManualTrigger }
            , effects
            )

        BuildJobDetailsFetched (Err _) ->
            -- https://github.com/concourse/concourse/issues/3201
            ( model, effects )

        _ ->
            ( model, effects )


handleDelivery : Delivery -> ET Model
handleDelivery delivery ( model, effects ) =
    case delivery of
        KeyDown keyEvent ->
            handleKeyPressed keyEvent ( model, effects )

        KeyUp keyEvent ->
            case keyEvent.code of
                Keyboard.T ->
                    ( { model | previousTriggerBuildByKey = False }, effects )

                _ ->
                    ( model, effects )

        ClockTicked OneSecond time ->
            let
                newModel =
                    { model
                        | now = Just time
                        , hoveredCounter = model.hoveredCounter + 1
                    }
            in
            updateOutput
                (Build.Output.Output.handleStepTreeMsg <|
                    StepTree.updateTooltip newModel
                )
                ( newModel, effects )

        ScrolledToBottom atBottom ->
            ( { model | autoScroll = atBottom }, effects )

        EventsReceived result ->
            let
                scrollEffects =
                    case getScrollBehavior model of
                        ScrollWindow ->
                            effects ++ [ Effects.Scroll Effects.ToBottom ]

                        NoScroll ->
                            effects
            in
            case result of
                Ok envelopes ->
                    let
                        newModel =
                            if
                                List.map .data envelopes
                                    |> List.member STModels.NetworkError
                            then
                                { model | authorized = False }

                            else
                                model
                    in
                    updateOutput
                        (Build.Output.Output.handleEnvelopes envelopes)
                        ( newModel, scrollEffects )

                Err _ ->
                    let
                        eventSourceOpened =
                            model.currentBuild
                                |> RemoteData.toMaybe
                                |> Maybe.andThen .output
                                |> Maybe.map .eventSourceOpened
                                |> Maybe.withDefault True
                    in
                    ( if eventSourceOpened then
                        -- connection could have dropped out of the blue;
                        -- just let the browser handle reconnecting
                        model

                      else
                        -- assume request was rejected because auth is required;
                        -- no way to really tell
                        { model | authorized = False }
                    , scrollEffects
                    )

        ElementVisible ( id, True ) ->
            let
                lastBuildVisible =
                    model.history
                        |> List.Extra.last
                        |> Maybe.map .id
                        |> Maybe.map String.fromInt
                        |> Maybe.map ((==) id)
                        |> Maybe.withDefault False

                hasNextPage =
                    model.nextPage /= Nothing

                needsToFetchMorePages =
                    not model.fetchingHistory && lastBuildVisible && hasNextPage
            in
            case currentJob model of
                Just job ->
                    if needsToFetchMorePages then
                        ( { model | fetchingHistory = True }
                        , effects ++ [ FetchBuildHistory job model.nextPage ]
                        )

                    else
                        ( model, effects )

                Nothing ->
                    ( model, effects )

        ElementVisible ( id, False ) ->
            let
                currentBuildInvisible =
                    model.currentBuild
                        |> RemoteData.toMaybe
                        |> Maybe.map (.build >> .id >> String.fromInt)
                        |> Maybe.map ((==) id)
                        |> Maybe.withDefault False

                shouldScroll =
                    currentBuildInvisible && not model.scrolledToCurrentBuild
            in
            ( { model | scrolledToCurrentBuild = True }
            , effects
                ++ (if shouldScroll then
                        [ Scroll <| ToId id ]

                    else
                        []
                   )
            )

        _ ->
            ( model, effects )


update : Message -> ET Model
update msg ( model, effects ) =
    case msg of
        SwitchToBuild build ->
            ( model
            , effects
                ++ [ NavigateTo <| Routes.toString <| Routes.buildRoute build ]
            )

        Hover state ->
            let
                newModel =
                    { model | hoveredElement = state, hoveredCounter = 0 }
            in
            updateOutput
                (Build.Output.Output.handleStepTreeMsg <| StepTree.updateTooltip newModel)
                ( newModel, effects )

        TriggerBuild job ->
            case job of
                Nothing ->
                    ( model, effects )

                Just someJob ->
                    ( model, effects ++ [ DoTriggerBuild someJob ] )

        AbortBuild buildId ->
            ( model, effects ++ [ DoAbortBuild buildId ] )

        ToggleStep id ->
            updateOutput
                (Build.Output.Output.handleStepTreeMsg <| StepTree.toggleStep id)
                ( model, effects )

        SwitchTab id tab ->
            updateOutput
                (Build.Output.Output.handleStepTreeMsg <| StepTree.switchTab id tab)
                ( model, effects )

        SetHighlight id line ->
            updateOutput
                (Build.Output.Output.handleStepTreeMsg <| StepTree.setHighlight id line)
                ( model, effects )

        ExtendHighlight id line ->
            updateOutput
                (Build.Output.Output.handleStepTreeMsg <| StepTree.extendHighlight id line)
                ( model, effects )

        ScrollBuilds event ->
            let
                scroll =
                    if event.deltaX == 0 then
                        [ Scroll (Element "builds" event.deltaY) ]

                    else
                        [ Scroll (Element "builds" -event.deltaX) ]

                checkVisibility =
                    case model.history |> List.Extra.last of
                        Just build ->
                            [ Effects.CheckIsVisible <| String.fromInt build.id ]

                        Nothing ->
                            []
            in
            ( model, effects ++ scroll ++ checkVisibility )

        GoToRoute route ->
            ( model, effects ++ [ NavigateTo <| Routes.toString <| route ] )

        _ ->
            ( model, effects )


getScrollBehavior : Model -> ScrollBehavior
getScrollBehavior model =
    if model.autoScroll then
        case model.currentBuild |> RemoteData.toMaybe of
            Nothing ->
                NoScroll

            Just cb ->
                case cb.build.status of
                    Concourse.BuildStatusSucceeded ->
                        NoScroll

                    Concourse.BuildStatusPending ->
                        NoScroll

                    _ ->
                        ScrollWindow

    else
        NoScroll


updateOutput :
    (OutputModel -> ( OutputModel, List Effect, Build.Output.Output.OutMsg ))
    -> ET Model
updateOutput updater ( model, effects ) =
    let
        currentBuild =
            model.currentBuild |> RemoteData.toMaybe
    in
    case ( currentBuild, currentBuild |> Maybe.andThen .output ) of
        ( Just cb, Just output ) ->
            let
                ( newOutput, outputEffects, outMsg ) =
                    updater output
            in
            handleOutMsg outMsg
                ( { model | currentBuild = RemoteData.Success { cb | output = Just newOutput } }
                , effects ++ outputEffects
                )

        _ ->
            ( model, effects )


handleKeyPressed : Keyboard.KeyEvent -> ET Model
handleKeyPressed keyEvent ( model, effects ) =
    let
        currentBuild =
            model.currentBuild |> RemoteData.toMaybe |> Maybe.map .build

        newModel =
            case ( model.previousKeyPress, keyEvent.shiftKey, keyEvent.code ) of
                ( Nothing, False, Keyboard.G ) ->
                    { model | previousKeyPress = Just keyEvent }

                _ ->
                    { model | previousKeyPress = Nothing }
    in
    if Keyboard.hasControlModifier keyEvent then
        ( newModel, [] )

    else
        case ( keyEvent.code, keyEvent.shiftKey ) of
            ( Keyboard.H, False ) ->
                case Maybe.andThen (nextBuild newModel.history) currentBuild of
                    Just build ->
                        update (SwitchToBuild build) ( newModel, effects )

                    Nothing ->
                        ( newModel, [] )

            ( Keyboard.L, False ) ->
                case
                    Maybe.andThen (prevBuild newModel.history) currentBuild
                of
                    Just build ->
                        update (SwitchToBuild build) ( newModel, effects )

                    Nothing ->
                        ( newModel, [] )

            ( Keyboard.J, False ) ->
                ( newModel, [ Scroll Down ] )

            ( Keyboard.K, False ) ->
                ( newModel, [ Scroll Up ] )

            ( Keyboard.T, True ) ->
                if not newModel.previousTriggerBuildByKey then
                    update
                        (TriggerBuild (currentBuild |> Maybe.andThen .job))
                        ( { newModel | previousTriggerBuildByKey = True }, effects )

                else
                    ( newModel, [] )

            ( Keyboard.A, True ) ->
                if currentBuild == List.head newModel.history then
                    case currentBuild of
                        Just build ->
                            update (AbortBuild build.id) ( newModel, effects )

                        Nothing ->
                            ( newModel, [] )

                else
                    ( newModel, [] )

            ( Keyboard.G, True ) ->
                ( { newModel | autoScroll = True }, [ Scroll ToBottom ] )

            ( Keyboard.G, False ) ->
                if
                    (model.previousKeyPress |> Maybe.map .code)
                        == Just Keyboard.G
                then
                    ( { newModel | autoScroll = False }, [ Scroll ToTop ] )

                else
                    ( newModel, [] )

            ( Keyboard.Slash, True ) ->
                ( { newModel | showHelp = not newModel.showHelp }, [] )

            _ ->
                ( newModel, [] )


nextBuild : List Concourse.Build -> Concourse.Build -> Maybe Concourse.Build
nextBuild builds build =
    case builds of
        first :: second :: rest ->
            if second == build then
                Just first

            else
                nextBuild (second :: rest) build

        _ ->
            Nothing


prevBuild : List Concourse.Build -> Concourse.Build -> Maybe Concourse.Build
prevBuild builds build =
    case builds of
        first :: second :: rest ->
            if first == build then
                Just second

            else
                prevBuild (second :: rest) build

        _ ->
            Nothing


handleBuildFetched : Int -> Concourse.Build -> ET Model
handleBuildFetched browsingIndex build ( model, effects ) =
    if browsingIndex == model.browsingIndex then
        let
            currentBuild =
                case model.currentBuild |> RemoteData.toMaybe of
                    Nothing ->
                        { build = build
                        , prep = Nothing
                        , output = Nothing
                        }

                    Just cb ->
                        { cb | build = build }

            withBuild =
                { model
                    | currentBuild = RemoteData.Success currentBuild
                    , history = updateHistory build model.history
                }

            fetchJobAndHistory =
                case ( currentJob model, build.job ) of
                    ( Nothing, Just buildJob ) ->
                        [ FetchBuildJobDetails buildJob
                        , FetchBuildHistory buildJob Nothing
                        ]

                    _ ->
                        []

            ( newModel, cmd ) =
                if build.status == Concourse.BuildStatusPending then
                    ( withBuild, effects ++ pollUntilStarted browsingIndex build.id )

                else if build.reapTime == Nothing then
                    case
                        model.currentBuild
                            |> RemoteData.toMaybe
                            |> Maybe.andThen .prep
                    of
                        Nothing ->
                            initBuildOutput build ( withBuild, effects )

                        Just _ ->
                            let
                                ( newNewModel, newEffects ) =
                                    initBuildOutput build ( withBuild, effects )
                            in
                            ( newNewModel
                            , newEffects
                                ++ [ FetchBuildPrep
                                        1000
                                        browsingIndex
                                        build.id
                                   ]
                            )

                else
                    ( withBuild, effects )
        in
        ( { newModel | fetchingHistory = True }
        , cmd ++ fetchJobAndHistory ++ [ SetFavIcon (Just build.status) ]
        )

    else
        ( model, effects )


pollUntilStarted : Int -> Int -> List Effect
pollUntilStarted browsingIndex buildId =
    [ FetchBuild 1000 browsingIndex buildId
    , FetchBuildPrep 1000 browsingIndex buildId
    ]


initBuildOutput : Concourse.Build -> ET Model
initBuildOutput build ( model, effects ) =
    let
        ( output, outputCmd ) =
            Build.Output.Output.init model.highlight build
    in
    ( { model
        | currentBuild =
            RemoteData.map
                (\info -> { info | output = Just output })
                model.currentBuild
      }
    , effects ++ outputCmd
    )


handleHistoryFetched : Paginated Concourse.Build -> ET Model
handleHistoryFetched history ( model, effects ) =
    let
        currentBuild =
            model.currentBuild |> RemoteData.map .build

        newModel =
            { model
                | history = List.append model.history history.content
                , nextPage = history.pagination.nextPage
                , fetchingHistory = False
            }
    in
    case ( currentBuild, currentJob model ) of
        ( RemoteData.Success build, Just job ) ->
            if List.member build newModel.history then
                ( newModel, effects ++ [ CheckIsVisible <| String.fromInt build.id ] )

            else
                ( { newModel | fetchingHistory = True }, effects ++ [ FetchBuildHistory job history.pagination.nextPage ] )

        _ ->
            ( newModel, effects )


handleBuildPrepFetched : Int -> Concourse.BuildPrep -> ET Model
handleBuildPrepFetched browsingIndex buildPrep ( model, effects ) =
    if browsingIndex == model.browsingIndex then
        ( { model
            | currentBuild =
                RemoteData.map
                    (\info -> { info | prep = Just buildPrep })
                    model.currentBuild
          }
        , effects
        )

    else
        ( model, effects )


documentTitle : Model -> String
documentTitle =
    extractTitle


view : UserState -> Model -> Html Message
view userState model =
    Html.div
        (id "page-including-top-bar" :: Views.Styles.pageIncludingTopBar)
        [ Html.div
            (id "top-bar-app" :: Views.Styles.topBar False)
            [ TopBar.concourseLogo
            , TopBar.breadcrumbs <|
                case model.page of
                    OneOffBuildPage buildId ->
                        Routes.OneOffBuild
                            { id = buildId
                            , highlight = model.highlight
                            }

                    JobBuildPage buildId ->
                        Routes.Build
                            { id = buildId
                            , highlight = model.highlight
                            }
            , Login.view userState model False
            ]
        , Html.div
            (id "page-below-top-bar" :: Views.Styles.pipelinePageBelowTopBar)
            [ viewBuildPage model ]
        ]


viewBuildPage : Model -> Html Message
viewBuildPage model =
    case model.currentBuild |> RemoteData.toMaybe of
        Just currentBuild ->
            Html.div
                [ class "with-fixed-header"
                , attribute "data-build-name" currentBuild.build.name
                ]
                [ viewBuildHeader currentBuild.build model
                , body
                    { currentBuild = currentBuild
                    , authorized = model.authorized
                    , showHelp = model.showHelp
                    }
                ]

        _ ->
            LoadingIndicator.view


body :
    { currentBuild : CurrentBuild
    , authorized : Bool
    , showHelp : Bool
    }
    -> Html Message
body { currentBuild, authorized, showHelp } =
    Html.div [ class "scrollable-body build-body" ] <|
        if authorized then
            [ viewBuildPrep currentBuild.prep
            , Html.Lazy.lazy viewBuildOutput currentBuild.output
            , keyboardHelp showHelp
            ]
                ++ tombstone currentBuild

        else
            [ NotAuthorized.view ]


keyboardHelp : Bool -> Html Message
keyboardHelp showHelp =
    Html.div
        [ classList
            [ ( "keyboard-help", True )
            , ( "hidden", not showHelp )
            ]
        ]
        [ Html.div
            [ class "help-title" ]
            [ Html.text "keyboard shortcuts" ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "h" ]
                , Html.span [ class "key" ] [ Html.text "l" ]
                ]
            , Html.text "previous/next build"
            ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "j" ]
                , Html.span [ class "key" ] [ Html.text "k" ]
                ]
            , Html.text "scroll down/up"
            ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "T" ] ]
            , Html.text "trigger a new build"
            ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "A" ] ]
            , Html.text "abort build"
            ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "gg" ] ]
            , Html.text "scroll to the top"
            ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "G" ] ]
            , Html.text "scroll to the bottom"
            ]
        , Html.div
            [ class "help-line" ]
            [ Html.div
                [ class "keys" ]
                [ Html.span [ class "key" ] [ Html.text "?" ] ]
            , Html.text "hide/show help"
            ]
        ]


tombstone : CurrentBuild -> List (Html Message)
tombstone currentBuild =
    let
        build =
            currentBuild.build

        maybeBirthDate =
            Maybe.Extra.or build.duration.startedAt build.duration.finishedAt
    in
    case ( maybeBirthDate, build.reapTime ) of
        ( Just birthDate, Just reapTime ) ->
            [ Html.div
                [ class "tombstone" ]
                [ Html.div [ class "heading" ] [ Html.text "RIP" ]
                , Html.div
                    [ class "job-name" ]
                    [ Html.text <|
                        Maybe.withDefault
                            "one-off build"
                        <|
                            Maybe.map .jobName build.job
                    ]
                , Html.div
                    [ class "build-name" ]
                    [ Html.text <|
                        "build #"
                            ++ (case build.job of
                                    Nothing ->
                                        String.fromInt build.id

                                    Just _ ->
                                        build.name
                               )
                    ]
                , Html.div
                    [ class "date" ]
                    [ Html.text <|
                        mmDDYY birthDate
                            ++ "-"
                            ++ mmDDYY reapTime
                    ]
                , Html.div
                    [ class "epitaph" ]
                    [ Html.text <|
                        case build.status of
                            Concourse.BuildStatusSucceeded ->
                                "It passed, and now it has passed on."

                            Concourse.BuildStatusFailed ->
                                "It failed, and now has been forgotten."

                            Concourse.BuildStatusErrored ->
                                "It errored, but has found forgiveness."

                            Concourse.BuildStatusAborted ->
                                "It was never given a chance."

                            _ ->
                                "I'm not dead yet."
                    ]
                ]
            , Html.div
                [ class "explanation" ]
                [ Html.text "This log has been "
                , Html.a
                    [ Html.Attributes.href "https://concourse-ci.org/jobs.html#job-build-logs-to-retain" ]
                    [ Html.text "reaped." ]
                ]
            ]

        _ ->
            []


mmDDYY : Time.Posix -> String
mmDDYY =
    DateFormat.format
        [ DateFormat.monthFixed
        , DateFormat.text "/"
        , DateFormat.dayOfMonthFixed
        , DateFormat.text "/"
        , DateFormat.yearNumberLastTwo
        ]
        -- https://github.com/concourse/concourse/issues/2226
        Time.utc


viewBuildOutput : Maybe OutputModel -> Html Message
viewBuildOutput output =
    case output of
        Just o ->
            Build.Output.Output.view o

        Nothing ->
            Html.div [] []


viewBuildPrep : Maybe Concourse.BuildPrep -> Html Message
viewBuildPrep buildPrep =
    case buildPrep of
        Just prep ->
            Html.div [ class "build-step" ]
                [ Html.div
                    [ class "header"
                    , style "display" "flex"
                    , style "align-items" "center"
                    ]
                    [ Icon.icon
                        { sizePx = 15, image = "ic-cogs.svg" }
                        [ style "margin" "6.5px"
                        , style "margin-right" "0.5px"
                        , style "background-size" "contain"
                        ]
                    , Html.h3 [] [ Html.text "preparing build" ]
                    ]
                , Html.div []
                    [ Html.ul [ class "prep-status-list" ]
                        ([ viewBuildPrepLi "checking pipeline is not paused" prep.pausedPipeline Dict.empty
                         , viewBuildPrepLi "checking job is not paused" prep.pausedJob Dict.empty
                         ]
                            ++ viewBuildPrepInputs prep.inputs
                            ++ [ viewBuildPrepLi "waiting for a suitable set of input versions" prep.inputsSatisfied prep.missingInputReasons
                               , viewBuildPrepLi "checking max-in-flight is not reached" prep.maxRunningBuilds Dict.empty
                               ]
                        )
                    ]
                ]

        Nothing ->
            Html.div [] []


viewBuildPrepInputs : Dict String Concourse.BuildPrepStatus -> List (Html Message)
viewBuildPrepInputs inputs =
    List.map viewBuildPrepInput (Dict.toList inputs)


viewBuildPrepInput : ( String, Concourse.BuildPrepStatus ) -> Html Message
viewBuildPrepInput ( name, status ) =
    viewBuildPrepLi ("discovering any new versions of " ++ name) status Dict.empty


viewBuildPrepDetails : Dict String String -> Html Message
viewBuildPrepDetails details =
    Html.ul [ class "details" ]
        (List.map viewDetailItem (Dict.toList details))


viewDetailItem : ( String, String ) -> Html Message
viewDetailItem ( name, status ) =
    Html.li []
        [ Html.text (name ++ " - " ++ status) ]


viewBuildPrepLi : String -> Concourse.BuildPrepStatus -> Dict String String -> Html Message
viewBuildPrepLi text status details =
    Html.li
        [ classList
            [ ( "prep-status", True )
            , ( "inactive", status == Concourse.BuildPrepStatusUnknown )
            ]
        ]
        [ Html.div
            [ style "align-items" "center"
            , style "display" "flex"
            ]
            [ viewBuildPrepStatus status
            , Html.span []
                [ Html.text text ]
            ]
        , viewBuildPrepDetails details
        ]


viewBuildPrepStatus : Concourse.BuildPrepStatus -> Html Message
viewBuildPrepStatus status =
    case status of
        Concourse.BuildPrepStatusUnknown ->
            Html.div
                [ title "thinking..." ]
                [ Spinner.spinner { size = "12px", margin = "0 5px 0 0" } ]

        Concourse.BuildPrepStatusBlocking ->
            Html.div
                [ title "blocking" ]
                [ Spinner.spinner { size = "12px", margin = "0 5px 0 0" } ]

        Concourse.BuildPrepStatusNotBlocking ->
            Icon.icon
                { sizePx = 12
                , image = "ic-not-blocking-check.svg"
                }
                [ style "margin-right" "5px"
                , style "background-size" "contain"
                , title "not blocking"
                ]


viewBuildHeader : Concourse.Build -> Model -> Html Message
viewBuildHeader build model =
    let
        triggerButton =
            case currentJob model of
                Just _ ->
                    let
                        buttonDisabled =
                            model.disableManualTrigger

                        buttonHovered =
                            model.hoveredElement == Just TriggerBuildButton
                    in
                    Html.button
                        ([ attribute "role" "button"
                         , attribute "tabindex" "0"
                         , attribute "aria-label" "Trigger Build"
                         , attribute "title" "Trigger Build"
                         , onLeftClick <| TriggerBuild build.job
                         , onMouseEnter <| Hover <| Just TriggerBuildButton
                         , onFocus <| Hover <| Just TriggerBuildButton
                         , onMouseLeave <| Hover Nothing
                         , onBlur <| Hover Nothing
                         ]
                            ++ Styles.triggerButton
                                buttonDisabled
                                buttonHovered
                                build.status
                        )
                    <|
                        [ Icon.icon
                            { sizePx = 40
                            , image = "ic-add-circle-outline-white.svg"
                            }
                            []
                        ]
                            ++ (if buttonDisabled && buttonHovered then
                                    [ Html.div
                                        Styles.triggerTooltip
                                        [ Html.text <|
                                            "manual triggering disabled "
                                                ++ "in job config"
                                        ]
                                    ]

                                else
                                    []
                               )

                Nothing ->
                    Html.text ""

        abortHovered =
            model.hoveredElement == Just AbortBuildButton

        abortButton =
            if Concourse.BuildStatus.isRunning build.status then
                Html.button
                    ([ onLeftClick (AbortBuild build.id)
                     , attribute "role" "button"
                     , attribute "tabindex" "0"
                     , attribute "aria-label" "Abort Build"
                     , attribute "title" "Abort Build"
                     , onMouseEnter <| Hover <| Just AbortBuildButton
                     , onFocus <| Hover <| Just AbortBuildButton
                     , onMouseLeave <| Hover Nothing
                     , onBlur <| Hover Nothing
                     ]
                        ++ Styles.abortButton abortHovered
                    )
                    [ Icon.icon
                        { sizePx = 40
                        , image = "ic-abort-circle-outline-white.svg"
                        }
                        []
                    ]

            else
                Html.text ""

        buildTitle =
            case build.job of
                Just jobId ->
                    let
                        jobRoute =
                            Routes.Job { id = jobId, page = Nothing }
                    in
                    Html.a
                        [ StrictEvents.onLeftClick <| GoToRoute jobRoute
                        , href <| Routes.toString jobRoute
                        ]
                        [ Html.span [ class "build-name" ] [ Html.text jobId.jobName ]
                        , Html.text (" #" ++ build.name)
                        ]

                _ ->
                    Html.text ("build #" ++ String.fromInt build.id)
    in
    Html.div [ class "fixed-header" ]
        [ Html.div
            ([ id "build-header"
             , class "build-header"
             ]
                ++ Styles.header build.status
            )
            [ Html.div []
                [ Html.h1 [] [ buildTitle ]
                , case model.now of
                    Just n ->
                        BuildDuration.view build.duration n

                    Nothing ->
                        Html.text ""
                ]
            , Html.div
                [ style "display" "flex" ]
                [ abortButton, triggerButton ]
            ]
        , Html.div
            [ onMouseWheel ScrollBuilds ]
            [ lazyViewHistory build model.history ]
        ]


lazyViewHistory : Concourse.Build -> List Concourse.Build -> Html Message
lazyViewHistory currentBuild builds =
    Html.Lazy.lazy2 viewHistory currentBuild builds


viewHistory : Concourse.Build -> List Concourse.Build -> Html Message
viewHistory currentBuild builds =
    Html.ul [ id "builds" ]
        (List.map (viewHistoryItem currentBuild) builds)


viewHistoryItem : Concourse.Build -> Concourse.Build -> Html Message
viewHistoryItem currentBuild build =
    Html.li
        ([ classList [ ( "current", build.id == currentBuild.id ) ]
         , id <| String.fromInt build.id
         ]
            ++ Styles.historyItem build.status
        )
        [ Html.a
            [ onLeftClick <| SwitchToBuild build
            , href <| Routes.toString <| Routes.buildRoute build
            ]
            [ Html.text build.name ]
        ]


handleOutMsg : Build.Output.Output.OutMsg -> ET Model
handleOutMsg outMsg ( model, effects ) =
    case outMsg of
        Build.Output.Output.OutNoop ->
            ( model, effects )

        Build.Output.Output.OutBuildStatus status date ->
            case model.currentBuild |> RemoteData.toMaybe of
                Nothing ->
                    ( model, effects )

                Just currentBuild ->
                    let
                        build =
                            currentBuild.build

                        duration =
                            build.duration

                        newDuration =
                            if Concourse.BuildStatus.isRunning status then
                                duration

                            else
                                { duration | finishedAt = Just date }

                        newStatus =
                            if Concourse.BuildStatus.isRunning build.status then
                                status

                            else
                                build.status

                        newBuild =
                            { build | status = newStatus, duration = newDuration }
                    in
                    ( { model
                        | history = updateHistory newBuild model.history
                        , currentBuild = RemoteData.Success { currentBuild | build = newBuild }
                      }
                    , if Concourse.BuildStatus.isRunning build.status then
                        effects ++ [ SetFavIcon (Just status) ]

                      else
                        effects
                    )


updateHistory : Concourse.Build -> List Concourse.Build -> List Concourse.Build
updateHistory newBuild =
    List.map <|
        \build ->
            if build.id == newBuild.id then
                newBuild

            else
                build
