describe('Visor maps todo list', function() {
    // //WARNING, en caso de fallar este test debe borrar el usuario creado test@e2e.cl
    var URLBASE = 'http://20.114.239.7:8080'; //cambiar por URL front
    it('- 1. Create map', async function() {
        browser.waitForAngularEnabled(false)
        browser.get(URLBASE+'/#/inicio')
        
        browser.sleep(2000)
        // Login
        var textBoxUser = element(by.id('username'))
        var textBoxPass = element(by.id('password'))
        var btnLogin = element(by.xpath('/html/body/body/main/div/div[4]/button'))
        textBoxUser.sendKeys('admin@visor.cl')
        textBoxPass.sendKeys('holahola')
        btnLogin.click()

        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/dashboard')

        // Maps panel
        var btnMaps = element(by.xpath('/html/body/body/header/nav/a[3]'))
        btnMaps.click()

        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/mapas')

        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).isPresent()).toBe(true)
        
        
        var btnNewMap = element(by.xpath('/html/body/body/main/div[3]/button'))
        btnNewMap.click()
        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/mapas/nuevo-mapa')
        
        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).getText()).toContain('Nuevo mapa')

        //Create Map
        // var MapName
        element(by.xpath('/html/body/body/main/div[3]/input[1]')).sendKeys('Test map name')
        // var DescName
        element(by.xpath('/html/body/body/main/div[3]/textarea')).sendKeys('Descripcion test map')
        // var Imageurl
        element(by.xpath('/html/body/body/main/div[3]/input[2]')).sendKeys('URL test map')
        browser.sleep(3000)
        
        // create maps
        element(by.xpath('/html/body/body/main/div[3]/div/button[2]')).click()
        browser.sleep(2000)

        //Check Map created
        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).getText()).toContain('Test map name')
        expect(element(by.xpath('/html/body/body/main/div[2]/p[2]')).getText()).toContain('Descripcion test map')
    });

    it('- 2. Delete map', async function() {
        // Delete Map created
        var btnDeleteMaps = element(by.xpath('/html/body/body/main/div[3]/button[7]'))
        btnDeleteMaps.click()
        browser.sleep(2000)
        var btnConfirmDeleteMaps =element(by.xpath('/html/body/body/main/div[2]/button[1]'))
        btnConfirmDeleteMaps.click()
        browser.sleep(1000)

        browser.get(URLBASE+'/#/mapas')
        browser.sleep(3000)
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).getText()).not.toContain('Test map name')
    });

    it('-3. Navegate and ReviewAttribute table', async function() {
        // Maps panel
        var btnMaps = element(by.xpath('/html/body/body/header/nav/a[3]'))
        btnMaps.click()

        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/mapas')

        // Select first map
        element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).click()
        // Open map
        element(by.xpath('/html/body/body/main/div[3]/button[1]')).click()
        browser.sleep(1000)
        expect(element(by.xpath('/html/body/body/main/div/div[3]/div[2]/div[2]/div[3]/div[5]/div')).getText()).toContain('1000 km')
        // Zoom in
        element(by.xpath('/html/body/body/main/div/div[3]/div[2]/div[2]/div[3]/div[1]/button[1]')).click()
        browser.sleep(1000)
        expect(element(by.xpath('/html/body/body/main/div/div[3]/div[2]/div[2]/div[3]/div[5]/div')).getText()).toContain('500 km')
        // Zoom out
        element(by.xpath('/html/body/body/main/div/div[3]/div[2]/div[2]/div[3]/div[1]/button[2]')).click()
        browser.sleep(1000)
        element(by.xpath('/html/body/body/main/div/div[3]/div[2]/div[2]/div[3]/div[1]/button[2]')).click()
        browser.sleep(1000)
        expect(element(by.xpath('/html/body/body/main/div/div[3]/div[2]/div[2]/div[3]/div[5]/div')).getText()).toContain('2000 km')

        // Logout
        browser.get(URLBASE+'/#/mapas')
        browser.sleep(2000)
        var btnAccount = element(by.xpath('//*[@id="header"]/div[2]/a'))
        var btnLogout = element(by.xpath('//*[@id="header"]/div[2]/div/a[2]'))
        btnAccount.click()
        btnLogout.click()

        browser.sleep(2000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/inicio')
    });
});