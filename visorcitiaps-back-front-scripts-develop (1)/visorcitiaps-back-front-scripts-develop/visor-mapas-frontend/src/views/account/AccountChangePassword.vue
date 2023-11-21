<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Cambiar contraseña</h2>
    </div>

    <div class="container-body">

        <label for="old_password">Contraseña actual</label>
        <input id="old_password" v-model="account.old_password" type="password"
        name="contraseña actual" :class="{'error' : errors.has('contraseña actual') }" v-validate="'required'"
        :placeholder="[[errors.has('contraseña actual') ? errors.first('contraseña actual') : 'Ingresa la contraseña actual']]"/>

        <label for="new_password">Contraseña nueva</label>
        <input id="new_password" v-model="account.new_password" type="password" ref="contraseña nueva"
        name="contraseña nueva" :class="{'error' : errors.has('contraseña nueva') }" v-validate="'required'"
        :placeholder="[[errors.has('contraseña nueva') ? errors.first('contraseña nueva') : 'Ingresa la contraseña nueva']]"/>

        <label for="repeat_password">Repita la contraseña nueva</label>
        <input id="repeat_password" v-model="repeat_password" type="password"
        name="contraseña repetida" :class="{'error' : errors.has('contraseña repetida') }" v-validate="'required|confirmed:contraseña nueva'"  data-vv-as="contraseña nueva"
        :placeholder="[[errors.has('contraseña repetida') ? errors.first('contraseña repetida') : 'Ingresa la contraseña repetida']]"/>

        <span>{{errors.first('contraseña repetida')}}</span>

        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Cambiar contraseña</button>
        </div>

    </div>

</main>
</template>

<style>

</style>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

var moment = require('moment');

export default{
    name: 'test',
    data(){
        return{
            myself: null,
            account: {
                id: null,
                new_password: null,
                old_password: null
            },
            repeat_password: null
        }
    },
    components:{
        Breadcrumb
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Cuenta', path: '/cuenta' },
              {name:'Cambiar contraseña'}
              ]
        },
        routesNav: function(){
          return [
                {name:'Cuenta'},
                {name:'Cambiar contraseña' , path: '/cuenta/cambiar-contrasena'},
              ]
        },
    },
    props: [],
    methods:{
        dateFormatter: function (date) {
            moment.locale('es');
            return 'Creado el ' + moment(date).format('LL');
        },
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/cuenta')
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    console.log('OK')
                    that.changePassword();
                } else {
                    console.log('NOPE')
                }
            })
        },
        getMyself: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/accounts/me')
            .then(function (response) {
                that.myself = response.data;
                that.account.id = response.data.id;
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        changePassword: function(){
            var that = this;
            this.$http.put(this.$store.getters.backendurl + 'api/v1/accounts/password',
            this.account
            )
            .then(function (response) {
                that.$router.push('/cuenta')
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions? true:this.$router.push('/inicio')
        this.$store.getters.getPermissions.visor? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getMyself();
    }
}
</script>
