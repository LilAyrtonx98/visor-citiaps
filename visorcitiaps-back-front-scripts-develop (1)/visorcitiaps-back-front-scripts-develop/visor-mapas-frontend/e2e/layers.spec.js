describe('Visor Layer todo list', function() {
    // //WARNING, en caso de fallar este test debe borrar el usuario creado test@e2e.cl
    var URLBASE = 'http://20.114.239.7:8080'; //cambiar por URL front
    it('- 1. Create layer OWS', async function() {
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

        // Layer panel
        var btnLayer = element(by.xpath('/html/body/body/header/nav/a[4]'))
        btnLayer.click()

        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas')

        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).isPresent()).toBe(true)
        
        var btnNewLayer = element(by.xpath('/html/body/body/main/div[3]/button[1]'))
        btnNewLayer.click()
        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas/nueva-capa')
        expect(element(by.xpath('/html/body/body/main/div[3]/label[1]')).getText()).toContain('Nombre')

        //Create Layer
        // var LayerName
        element(by.xpath('/html/body/body/main/div[3]/input[1]')).sendKeys('LayerTestOWS')
        // select category
        element(by.xpath('/html/body/body/main/div[3]/select/option[2]')).click()
        // var textBoxLastName
        element(by.xpath('/html/body/body/main/div[3]/textarea')).sendKeys('Descripcion test ows de la capa')
        // OWS option
        element(by.xpath('/html/body/body/main/div[3]/div[1]/input[2]')).click()
        browser.sleep(3000)
        // URL service
        element(by.xpath('/html/body/body/main/div[3]/input[2]')).sendKeys('http://20.114.239.7:8600/geoserver/topp/ows?service=WFS&version=1.0.0&request=GetFeature&typeName=topp%3Astates&maxFeatures=50')
        // workspace
        element(by.xpath('/html/body/body/main/div[3]/input[3]')).sendKeys('topp')
        // datastore
        element(by.xpath('/html/body/body/main/div[3]/input[4]')).sendKeys('states')
        
        // create layer
        element(by.xpath('/html/body/body/main/div[3]/div[2]/button[2]')).click()
        browser.sleep(2000)
        // assign one map
        element(by.xpath('/html/body/body/main/div[3]/div[2]/div[1]/div/span[3]/button')).click()
        // End create layer
        element(by.xpath('/html/body/body/main/div[4]/div[1]/span[2]/button')).click()
        browser.sleep(2000)

        //Check layer created
        expect(element(by.xpath('/html/body/body/main/div[2]/p[3]')).getText()).toContain('Descripcion test ows de la capa')
        expect(element(by.xpath('/html/body/body/main/div[2]/p[6]/span')).getText()).toContain('topp')
        expect(element(by.xpath('/html/body/body/main/div[2]/p[7]/span')).getText()).toContain('states')
        
    });

    it('- 2. Delete OWS layer', async function() {
        // Delete layer created
        var btnDeleteLayer = element(by.xpath('/html/body/body/main/div[3]/button[4]'))
        btnDeleteLayer.click()
        browser.sleep(2000)
        var btnConfirmDeleteLayer =element(by.xpath('/html/body/body/main/div[2]/button[1]'))
        btnConfirmDeleteLayer.click()
        browser.sleep(1000)

        browser.get(URLBASE+'/#/capas')
        browser.sleep(3000)
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).getText()).not.toContain('LayerTestOWS')
        
    });

    it('- 3. Create layer ESRI', async function() {
        // Layer panel
        var btnLayer = element(by.xpath('/html/body/body/header/nav/a[4]'))
        btnLayer.click()

        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas')

        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).isPresent()).toBe(true)
        
        var btnNewLayer = element(by.xpath('/html/body/body/main/div[3]/button[1]'))
        btnNewLayer.click()
        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas/nueva-capa')
        expect(element(by.xpath('/html/body/body/main/div[3]/label[1]')).getText()).toContain('Nombre')

        //Create Layer
        // var LayerName
        element(by.xpath('/html/body/body/main/div[3]/input[1]')).sendKeys('LayerTestEsri')
        // select category
        element(by.xpath('/html/body/body/main/div[3]/select/option[2]')).click()
        // var textBoxLastName
        element(by.xpath('/html/body/body/main/div[3]/textarea')).sendKeys('Descripcion test esri de la capa')
        // ESRI option
        element(by.xpath('/html/body/body/main/div[3]/div[1]/input[3]')).click()
        browser.sleep(3000)
        // URL service
        element(by.xpath('/html/body/body/main/div[3]/input[2]')).sendKeys('https://geoportal.sag.gob.cl/server/rest/services/LimiteChile2019_line/MapServer/')
        
        // create layer
        element(by.xpath('/html/body/body/main/div[3]/div[2]/button[2]')).click()
        browser.sleep(2000)
        // assign one map
        element(by.xpath('/html/body/body/main/div[3]/div[2]/div[1]/div/span[3]/button')).click()
        // End create layer
        element(by.xpath('/html/body/body/main/div[4]/div[1]/span[2]/button')).click()
        browser.sleep(2000)

        //Check layer created
        expect(element(by.xpath('/html/body/body/main/div[2]/p[3]')).getText()).toContain('Descripcion test esri de la capa')
        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).getText()).toContain('LayerTestEsri')
    });

    it('- 4. Delete ESRI layer', async function() {
        // Delete layer created
        var btnDeleteLayer = element(by.xpath('/html/body/body/main/div[3]/button[4]'))
        btnDeleteLayer.click()
        browser.sleep(2000)
        var btnConfirmDeleteLayer =element(by.xpath('/html/body/body/main/div[2]/button[1]'))
        btnConfirmDeleteLayer.click()
        browser.sleep(1000)

        browser.get(URLBASE+'/#/capas')
        browser.sleep(3000)
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).getText()).not.toContain('LayerTestEsri')
    });

    it('- 5. Create layer SHP', async function() {
        // Layer panel
        var btnLayer = element(by.xpath('/html/body/body/header/nav/a[4]'))
        btnLayer.click()

        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas')

        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).isPresent()).toBe(true)
        
        var btnNewLayer = element(by.xpath('/html/body/body/main/div[3]/button[1]'))
        btnNewLayer.click()
        browser.sleep(3000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas/nueva-capa')
        expect(element(by.xpath('/html/body/body/main/div[3]/label[1]')).getText()).toContain('Nombre')

        //Create Layer
        // var LayerName
        element(by.xpath('/html/body/body/main/div[3]/input[1]')).sendKeys('LayerTestSHP')
        // select category
        element(by.xpath('/html/body/body/main/div[3]/select/option[2]')).click()
        // var textBoxLastName
        element(by.xpath('/html/body/body/main/div[3]/textarea')).sendKeys('Descripcion test esri de la capa')
        // SHP option
        element(by.xpath('/html/body/body/main/div[3]/div[1]/input[1]')).click()
        
        var uploadInput = element(by.css("input[type=file]"))
        //For local change the path, use fully not relative path
        uploadInput.sendKeys("/builds/cesar.kreep/visor-mapas-frontend/e2e/files-for-testing/nyc_roads.zip") 

        // take a breath 
        browser.driver.sleep(100)
        // Upload layer
        element(by.xpath('/html/body/body/main/div[3]/div[3]/button[2]')).click()
        browser.sleep(2000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/capas');

        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).getText()).toContain('Lista de capas')

        // Logout
        var btnAccount = element(by.xpath('//*[@id="header"]/div[2]/a'));
        var btnLogout = element(by.xpath('//*[@id="header"]/div[2]/div/a[2]'));
        btnAccount.click();
        btnLogout.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/inicio');
    });
});