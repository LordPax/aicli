# Changelog

## [Unreleased]

### Added

* Add parameter -u for text command to web page content

### Changed

* Parameter -f in text command accept image file (only for claude sdk)

## [0.4.0] - 2024-09-28

### Changed

* Parameter -s accepte now one value (breaking change)

## [0.3.0] - 2024-09-28

### Added

* Add parameter -i for text command to not make api call

### Changed

* Improve the way we initialize the sdk
* Rename `system` parameter to `context` for text command (breaking change)

## [0.2.0] - 2024-09-27

### Added

* Add parameter -L for text command to list all history name
* Add sdk mistral for text command

### Changed

* Improve the way we add message to history
* Improve input function for interactive mode

## [0.1.0] - 2024-09-22

### Added

* Add interactive mode for translate command
* Add sdk deepl for translate command
* Add interactive mode for text command
* Add parameter for text command
    * `sdk` : Select the sdk to use
    * `history` : Select the history to use
    * `model` : Select the model to use
    * `temp` : Set the temperature
    * `system` : Prompt with system role
    * `file` : Add content of file text to history
    * `clear` : Clear the history
    * `list-history` : List all history
* Add history with name for text command
* Add sdk claude for text command
* Add sdk openai for text command
* Add localized for `en` and `fr`
* Read config.ini file
