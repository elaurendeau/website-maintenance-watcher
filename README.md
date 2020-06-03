### Prerequisite
You must own an https://pushover.net account and install the mobile application.
The USER_KEY and TOKEN_KEY can be found on your profile page.

### Build
docker build -t website-maintenance-watcher .

### Run
docker run -e USER_KEY=XXXXXXXXXXXX -e TOKEN_KEY=XXXXXXXXXXXXXX website-maintenance-watcher:latest