describe('Visor Dashboard todo list', function() {
    var URLBASE = 'http://20.114.239.7:8080'; //cambiar por URL front
    it('- 1. Status dashboard page', async function() {
        browser.waitForAngularEnabled(false);
        browser.get(URLBASE+'/#/inicio');
        
        browser.sleep(3000);
        // Login
        var textBoxUser = element(by.id('username'));
        var textBoxPass = element(by.id('password'));
        var btnLogin = element(by.xpath('/html/body/body/main/div/div[4]/button'));
        textBoxUser.sendKeys('admin@visor.cl');
        textBoxPass.sendKeys('holahola');
        btnLogin.click();
        
        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/dashboard');
        
        // Dashboard resume status
        // Layers 
        expect(element(by.xpath('/html/body/body/main/div[1]/div[5]/span[1]')).isPresent()).toBe(true);
        // LayersCateg 
        expect(element(by.xpath('/html/body/body/main/div[1]/div[2]/span[1]')).isPresent()).toBe(true);
        // Users
        expect(element(by.xpath('/html/body/body/main/div[1]/div[6]/span[1]')).isPresent()).toBe(true);
        // Groups
        expect(element(by.xpath('/html/body/body/main/div[1]/div[3]/span[1]')).isPresent()).toBe(true);
        // Maps
        expect(element(by.xpath('/html/body/body/main/div[1]/div[1]/span[1]')).isPresent()).toBe(true);
        // Geo
        expect(element(by.xpath('/html/body/body/main/div[1]/div[4]/span[1]')).isPresent()).toBe(true);
        // MapsUser
        expect(element(by.xpath('/html/body/body/main/div[2]/div/span[1]')).isPresent()).toBe(true);
    });
    it('- 2. Status dashboard buttons functionality', async function() {
        // Dashboard buttons functionality
        // Help
        var btnHelp = element(by.xpath('/html/body/body/footer/div[1]/a'));
        btnHelp.click();
        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/ayuda');
        expect(element(by.xpath('/html/body/body/main/h3[6]')).getText()).toContain('Atajos de teclado');
        // Visor (status and check the first map)
        var btnVisor = element(by.xpath('/html/body/body/header/nav/a[2]'));
        btnVisor.click();
        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/visor');
        expect(element(by.xpath('/html/body/body/main/div/a[1]/div/div/p')).isPresent()).toBe(true);
        // Maps
        var btnMaps = element(by.xpath('/html/body/body/header/nav/a[3]'));
        btnMaps.click();
        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/mapas');
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).isPresent()).toBe(true);
        // Layers
        var btnLayers = element(by.xpath('/html/body/body/header/nav/a[4]'));
        btnLayers.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas');
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).isPresent()).toBe(true);
        // Geo
        var btnGeo = element(by.xpath('/html/body/body/header/nav/a[5]'));
        btnGeo.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/geoprocesos');
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).isPresent()).toBe(true);
        // User
        var btnUsers = element(by.xpath('/html/body/body/header/nav/a[6]'));
        btnUsers.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/usuarios');
        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).isPresent()).toBe(true);    
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