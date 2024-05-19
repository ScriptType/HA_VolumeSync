# Volume Sync with Home Assistant

This project is a Go application that integrates with Home Assistant to synchronize media volume and control media playback through keyboard shortcuts.

<img src="https://github.com/ScriptType/HA_VolumeSync/blob/main/volumeSync.gif" alt="Demo Video" width="480" height="270">

## Features

- Synchronize media volume with Home Assistant
- Control media playback (Play, Pause, Next, Previous, Mute) using keyboard shortcuts
- Interactive CLI menu for starting the program and updating configuration

## Table of Contents

- [Installation](#installation)
  - [Manual Installation](#manual-installation)
  - [Using Pre-built Release](#using-pre-built-release)
- [Configuration](#configuration)
- [Usage](#usage)
- [Keyboard Shortcuts](#keyboard-shortcuts)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Manual Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/volume-sync-homeassistant.git
   cd volume-sync-homeassistant
   ```

2. **Install dependencies:**

   This project uses the `github.com/manifoldco/promptui` library for the CLI menu. You can install it using:

   ```sh
   go get github.com/manifoldco/promptui
   ```

3. **Build the application:**

   ```sh
   go build
   ```

### Using Pre-built Release

1. **Download the latest release:**

[Releases](https://github.com/ScriptType/HA_VolumeSync/releases)

2. **Run the application:**

   Navigate to the directory where you downloaded the go-executable and run:

   ```sh
   sudo ./HA_VolumeSync
   ```

## Configuration

When you run the application for the first time, it will prompt you to enter the necessary configuration details: Base URL, Auth Token, and Entity ID. This configuration will be saved to `config.json`.

You can also update the configuration through the interactive CLI menu.

## Usage

1. **Run the application:**

   ```sh
   ./volume-sync-homeassistant
   ```

2. **Navigate the CLI menu:**

   - **Start Program**: Starts the volume synchronization and media control functionalities.
   - **Update Configuration**: Allows you to update the Base URL, Auth Token, or Entity ID.
   - **Exit**: Exits the application.

## Keyboard Shortcuts

The following keyboard shortcuts are supported for media control:

- **Media Mute**: Mute or unmute
- **Media Volume Down/Up**: Decrease or increase volume

The Following Keyboard shortcuts will be supported for media control:

- **Media Previous**: Go to the previous track
- **Media Play/Pause**: Toggle play/pause
- **Media Next**: Go to the next track

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add new feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Create a new Pull Request

Please make sure your code follows the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
