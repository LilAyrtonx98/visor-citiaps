exports.config = {
    capabilities: {
        'browserName': 'chrome',
        chromeOptions: {
            args: [
                '--no-sandbox',
                '--headless',
                '--disable-gpu',
                '--disable-dev-shm-usage',
                '--window-size=1200,600',
                '--disable-infobars',
            ],
            prefs: {
                'credentials_enable_service': false,
                'profile': {
                    'password_manager_enabled': false,
                }
            }
        },
    },
    directConnect: true,
    specs: ['*.spec.js'],
};