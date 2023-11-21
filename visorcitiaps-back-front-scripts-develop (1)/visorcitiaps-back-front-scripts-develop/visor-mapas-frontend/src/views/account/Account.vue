<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2 v-if="myself">{{myself.firstname + ' ' + myself.lastname}}</h2>
    </div>

    <div class="container-body">
        <strong>Mi información</strong>
        <p v-if="myself">{{myself.username}}</p>
        <p v-if="myGroup">{{myGroup.name}}</p>
        <p v-if="myself">{{dateFormatter(myself.created_at)}}</p>
        <br>
    </div>

    <SideNav
    class="container-sidenav"
    v-bind:links="routesNav"
    >
    </SideNav>

</main>
</template>

<style>

</style>

<script>
// @ is an alias to /src
import SideNav from '@/components/SideNav.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'

import {formatter} from '@/components/mixin/formatter.js'

export default{
    name: 'test',
    mixins: [formatter],
    data(){
        return{
            myself: null,
            myGroup: null
        }
    },
    components:{
        Breadcrumb,
        SideNav
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Cuenta' }
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
        getMyself: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/accounts/me')
            .then(function (response) {
                that.myself = response.data;
                that.getGroup();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getGroup: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/' + this.myself.id_group)
            .then(function (response) {
                that.myGroup = response.data;
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
