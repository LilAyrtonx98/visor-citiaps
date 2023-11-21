<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Nuevo usuario</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="firstname">Nombre</label>
        <input id="firstname" v-model="user.firstname" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre']]"/>

        <label for="lastname">Apellidos</label>
        <input id="lastname" v-model="user.lastname" type="text"
        name="apellidos" :class="{'error' : errors.has('apellidos') }" v-validate="'required'"
        :placeholder="[[errors.has('apellidos') ? errors.first('apellidos') : 'Ingresa los apellidos']]"/>

        <label for="group">Grupo del usuario</label>
        <br>
        <select id="group"
        v-if="groups"
        v-model="user.id_group"
        name="grupo del usuario" :class="{'error' : errors.has('grupo del usuario') }" v-validate="groupsIncluded">
            <option
            value="-1"
            disabled
            >
                {{errors.has('grupo del usuario') ? 'Selecciona un grupo para el usuario' : 'Selecciona un grupo'}}
            </option>
            <option
            v-for="item in groups"
            v-bind:key="item.index"
            v-bind:value="item.id">
            {{item.name}}
            </option>
        </select>
        <span v-else>No hay grupos que seleccionar para el usuario</span>
        <br>

        <label for="username">Correo electrónico</label>
        <input id="username" v-model="user.username" type="text"
        name="nombre de usuario" :class="{'error' : errors.has('nombre de usuario') }" v-validate="'required|email|unique'"
        :placeholder="[[errors.has('nombre de usuario') ? errors.first('nombre de usuario') : 'Escribe el nombre de usuario']]"/>
        
        <label for="password">Contraseña</label>
        <input id="password" v-model="user.password" type="password"
        name="contraseña" :class="{'error' : errors.has('contraseña') }" v-validate="'required'"
        :placeholder="[[errors.has('contraseña') ? errors.first('contraseña') : 'Ingresa la contraseña']]"/>

        <label for="permissions">Permisos de usuario</label>
        <div id="permissions">
            <input type="checkbox" id="checkbox" v-model="user.permissions.users">
            <label for="checkbox">Usuarios</label>
            <br>
            <input type="checkbox" id="checkbox" v-model="user.permissions.layers">
            <label for="checkbox">Capas</label>
            <br>
            <input type="checkbox" id="checkbox" v-model="user.permissions.geo">
            <label for="checkbox">Geoprocesos</label>
            <br>
            <input type="checkbox" id="checkbox" v-model="user.permissions.maps">
            <label for="checkbox">Mapas</label>
            <br>
            <input type="checkbox" id="checkbox" v-model="user.permissions.visor">
            <label for="checkbox">Visor</label>
        </div>

        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Registrar usuario</button>
        </div>

    </div>

    <div class="container-list"> 
        
    </div>
 
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restUser} from '@/components/mixin/restUser.js'
import {restGroup} from '@/components/mixin/restGroup.js'

export default{
    name: 'test',
    mixins: [restUser, restGroup],
    data(){
        return{
            
        }
    },
    components:{
        Breadcrumb
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Usuarios', path: '/usuarios' },
              {name: 'Nuevo usuario'}
              ]
        },
        groupsIncluded: function(){
            let list = 'included:';
            for (let i = 0; i < this.groups.length; i++) {
                list = list + this.groups[i].id.toString() + ',';
            }
            return list;
        },
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/')
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    that.postUser();
                } else {
                    console.log('Error')
                }
            })
        },
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroupsNoPages();

        //Verificar username único en servidor
        const isUnique = (value) => {
            return this.$http.post(this.$store.getters.backendurl  + 'api/v1/accounts/check', {username: value})
            .then(function (response) {
                //Obtener ID de la tarea recién creada y pasar a la vista de asignar usuarios
                return {
                    valid: response.data.valid,
                    data: {
                        message: response.data.message
                    }
                };
            })
            .catch(function (error) {
                console.log(error);
                return {
                    valid: false,
                    data: {
                        message: "No se puede verificar"
                    }
                }
            });
        };
        this.$validator.extend('unique', {
            validate: isUnique,
            getMessage: (field, params, data) => {
                return data.message;
            }
        });
    },
    watch:{
        'user.permissions.maps': function(value){
            if (value) this.user.permissions.annotation=true
            else this.user.permissions.annotation=false
        }
    }
}
</script>

<style lang="scss" scoped>

</style>
