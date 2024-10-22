# creepycast_poc

creepycast_poc is a podcast server that serves as a proof of concept for creating dynamic podcasts. This project is not production-ready and includes a lot of hacks and bad implementations.

## Overview

The server uses `ffmpeg` to combine a number of MP3 files on the fly when an episode is downloaded. The files combined are:

1. **Current Time Greeting**: An MP3 file corresponding to the current hour of the day (0-23.mp3).
2. **Generic Content**: A generic MP3 file that serves as the main content of the podcast.
3. **Podcatcher Specific**: An MP3 file specific to the podcatcher being used.
4. **Operating System Specific**: An MP3 file specific to the operating system of the listener.

This project demonstrates how 'injection' or 'dynamic podcasts' could be created by combining different audio segments dynamically.

## Disclaimer

This project is purely for demonstration purposes and should not be used in a production environment.