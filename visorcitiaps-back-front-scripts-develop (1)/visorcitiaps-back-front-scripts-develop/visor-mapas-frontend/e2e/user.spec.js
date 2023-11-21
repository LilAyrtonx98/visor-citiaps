describe('Visor User todo list', function() {
    //WARNING, en caso de fallar este test debe borrar el usuario creado test@e2e.cl
    var URLBASE = 'http://20.114.239.7:8080'; //cambiar por URL front
    it('- 1. Create user and permission assignment', async function() {
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

        // User panel
        var btnUsers = element(by.xpath('/html/body/body/header/nav/a[6]'))
        btnUsers.click()

        browser.sleep(1000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/usuarios')
        expect(element(by.xpath('/html/body/body/main/div[1]/h2')).isPresent()).toBe(true)
        
        var btnNewUser = element(by.xpath('/html/body/body/main/div[3]/button[1]'))
        btnNewUser.click()
        browser.sleep(1000)
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/usuarios/nuevo-usuario')
        expect(element(by.xpath('/html/body/body/main/div[3]/label[1]')).getText()).toContain('Nombre')

        //Create user
        // var textBoxName
        element(by.xpath('/html/body/body/main/div[3]/input[1]')).sendKeys('TestName')
        // var textBoxLastName
        element(by.xpath('/html/body/body/main/div[3]/input[2]')).sendKeys('TestLastName')
        // var textBoxUserName
        element(by.xpath('/html/body/body/main/div[3]/input[3]')).sendKeys('test@e2e.cl')
        // var textBoxPassword
        element(by.xpath('/html/body/body/main/div[3]/input[4]')).sendKeys('hola1234')
        element(by.xpath('/html/body/body/main/div[3]/select/option[2]')).click()
        
        // Check permissions
        var checkUser = element(by.xpath('/html/body/body/main/div[3]/div[1]/input[1]'))
        checkUser.click()
        var checkLayer = element(by.xpath('/html/body/body/main/div[3]/div[1]/input[2]'))
        checkLayer.click()
        var checkGeo = element(by.xpath('/html/body/body/main/div[3]/div[1]/input[3]'))
        checkGeo.click()
        var checkMap = element(by.xpath('/html/body/body/main/div[3]/div[1]/input[4]'))
        checkMap.click()
        var checkVisor = element(by.xpath('/html/body/body/main/div[3]/div[1]/input[5]'))
        checkVisor.click()

        // Save user
        var btnSave = element(by.xpath('/html/body/body/main/div[3]/div[2]/button[2]'))
        btnSave.click()
        browser.sleep(3000)

        // Assign all maps
        var btnAssignAll = element(by.xpath('/html/body/body/main/div[3]/div[1]/span/button'))
        btnAssignAll.click()
        // End create user
        var btnRegisterFin = element(by.xpath('/html/body/body/main/div[4]/div[1]/span[2]/button'))
        btnRegisterFin.click()
        browser.sleep(2000)

        //Check user created
        expect(element(by.xpath('/html/body/body/main/div[2]/p[1]')).getText()).toContain('test@e2e.cl')
        expect(element(by.xpath('/html/body/body/main/div[2]/p[2]')).getText()).toContain('Grupo Principal')
        expect(element(by.xpath('/html/body/body/main/div[2]/p[1]')).getText()).toContain('test@e2e.cl')
        
    });

    it('- 2. Delete user', async function() {
        // Delete user created
        var btnDeleteUser = element(by.xpath('/html/body/body/main/div[3]/button[4]'))
        btnDeleteUser.click()
        browser.sleep(2000)
        
        var btnConfirmDeleteUser =element(by.xpath('/html/body/body/main/div[2]/button[1]'))
        btnConfirmDeleteUser.click()
        browser.sleep(2000)

        browser.get(URLBASE+'/#/usuarios')
        browser.sleep(2000)
        expect(element(by.xpath('/html/body/body/main/div[2]/div[1]/span[1]/a')).getText()).not.toContain('TestName TestLastName')

        // Logout
        var btnAccount = element(by.xpath('//*[@id="header"]/div[2]/a'));
        var btnLogout = element(by.xpath('//*[@id="header"]/div[2]/div/a[2]'));
        btnAccount.click();
        btnLogout.click();

        browser.sleep(2000);
        expect(browser.getCurrentUrl()).toEqual(URLBASE+'/#/inicio');
        
    });
});