<template>
<div class="center-box">
    
    <div class="center-aux">
        <p class="title-text">Bienvenido</p>
    </div>

    <div class="center-input">
        <label for="username">Nombre de usuario</label>
        <input id="username" v-model="username" type="text" v-on:keyup.enter="onClickGoNext"
        name="nombre de usuario" :class="{'error' : errors.has('nombre de usuario') }" v-validate="'required'"
        :placeholder="[['Escriba su nombre de usuario']]"
        />
        <p class="error-label" v-if="errors.has('nombre de usuario')">{{errors.first('nombre de usuario')}}</p>
    </div>
    <br>
    <div class="center-input">
        <label for="password">Contraseña</label>
        <input id="password" v-model="password" type="password" v-on:keyup.enter="onClickGoNext"
        name="contraseña" :class="{'error' : errors.has('contraseña') }" v-validate="'required'"
        :placeholder="[['Escriba su contraseña']]"
        />
        <p class="error-label" v-if="errors.has('contraseña')">{{errors.first('contraseña')}}</p>
    </div>

    <span class="center-input" style="color: red;" v-if="message">{{'Ocurrió un error al iniciar sesión'}}</span>

    <div class="center-aux">
        <LoadSpinner 
        v-if="showSpinner"
        >
        </LoadSpinner>
        <button v-else class="btn btn-primary" v-on:click="onClickGoNext()">Iniciar sesión</button>
    </div>
</div>
</template>

<style>

</style>

<script>
import LoadSpinner from '@/components/LoadSpinner.vue'

export default{
    name: 'login',
    data(){
        return{
            username: null,
            password: null,
            message: null,
            showSpinner: false,
        }
    },
    components:{
        LoadSpinner
    },
    computed: {

    },
    props: [],
    methods:{
        hello: function() {
            console.log('Enter')
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    console.log('OK')
                    that.onClickLogin();
                } else {
                    console.log('NOPE')
                }
            })
        },
        onClickLogin: function(){
            var that = this;
            this.showSpinner = true;
            this.$http.post(this.$store.getters.backendurl + 'api/v1/login',
            {
                username: this.username,
                password: this.password
            })
            .then(function (response) {
                that.message = response.data.message;
                //Crear cookie
                document.cookie = "session_public=" + response.headers.userpermissions;
                //document.cookie = "session_token=" + response.headers.token;
                //Guardar permisos de usuario
                that.$store.commit('setPermissions');
                //Guardar token
                //that.$store.commit('setToken');
                //Por defecto, se envía el token de sesión en cada consulta
                //that.$http.defaults.headers.common = {'Authorization': 'Bearer ' + that.$store.getters.getToken}
                //Inicio
                that.showSpinner = false;
                that.$router.push('/dashboard')
            })
            .catch(function (error) {
                console.log(error)
                that.showSpinner = false;
                that.message = error.response.data.message;
            });
        },
    },
    updated () {

    },
    created () {

    },
    mounted () {

    }
}
</script>

<style lang="scss" scoped>
.center-box {
    border: solid 2px #03440c;
    border-radius: 6px;
    min-width: 30%;
    background: white;
    box-shadow: 0px 0px 1em black;;
}
.title-text {
    font-size: 1.5em;
    font-weight: bold;
}
.center-aux {
    display: flex;
    justify-content: center;
    margin-top: 1em; 
    margin-bottom: 1em;
}
.center-input {
    margin: 0 5%;
}
button {
    margin: 1em 0;
    width: 90%;
}
</style>