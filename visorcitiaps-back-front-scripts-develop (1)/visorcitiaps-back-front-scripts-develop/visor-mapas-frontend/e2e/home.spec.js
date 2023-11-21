describe('Visor homepage todo list', function() {
    var URLBASE = 'http://20.114.239.7:8080';//cambiar por URL front

    it('- 1. Status home page', async function() {
        browser.waitForAngularEnabled(false);
        browser.get(URLBASE+'/#/inicio'); 

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/inicio'); 
        // Home-login status
        var titletextBoxUser = element(by.xpath('//*[@id="app"]/main/div/div[2]/label'));
        var titletextBoxPass = element(by.xpath('//*[@id="app"]/main/div/div[3]/label'));
        var btnLogin = element(by.xpath('/html/body/body/main/div/div[4]/button'));
        expect(titletextBoxUser.getText()).toEqual('Nombre de usuario');
        expect(titletextBoxPass.getText()).toEqual('Contraseña');
        expect(btnLogin.getText()).toEqual('Iniciar sesión');
        // About status
        var btnAbout = element(by.xpath('/html/body/body/header/nav/a[2]'));
        btnAbout.click();

        browser.sleep(2000);
        var titleAbout = element(by.xpath('//*[@id="app"]/main/div[1]/h1'));
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/acerca-de'); 
        expect(titleAbout.getText()).toContain('Visor de Mapas');
    });

    it('- 2. Login test (enter through login))', async function() {
        browser.waitForAngularEnabled(false);
        browser.get(URLBASE+'/#/inicio');
        
        browser.sleep(2000);
        var textBoxUser = element(by.id('username'));
        var textBoxPass = element(by.id('password'));
        var btnLogin = element(by.xpath('/html/body/body/main/div/div[4]/button'));

        textBoxUser.sendKeys('admin@visor.cl');
        textBoxPass.sendKeys('holahola');
        btnLogin.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/dashboard');
    });

    it('- 3. Logout test (enter through login))', async function() {
        browser.waitForAngularEnabled(false);
        browser.get(URLBASE+'/#/dashboard'); 
        
        browser.sleep(2000);
        var btnAccount = element(by.xpath('//*[@id="header"]/div[2]/a'));
        var btnLogout = element(by.xpath('//*[@id="header"]/div[2]/div/a[2]'));
        btnAccount.click();
        btnLogout.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/inicio');
    });
});