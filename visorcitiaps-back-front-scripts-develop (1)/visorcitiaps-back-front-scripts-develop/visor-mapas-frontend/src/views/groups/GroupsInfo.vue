<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{group.name}}</h2>
    </div>

    <div class="container-body">
        <strong>Información</strong>
        <p>{{dateFormatter(group.created_at)}}</p>
        <p>{{group.desc}}</p>
        <br>
        <strong>Usuarios que pertenecen a este grupo</strong>
        <div class="list-item-area">
            <template 
            v-for="(item) in users" >
                <router-link 
                v-bind:key="item.index" 
                class="list-item-name"
                :to="'/usuarios/' + item.id"
                >
                    {{item.firstname + ' ' + item.lastname}}
                </router-link>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.username}}</span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>
        </div>

        <Pagination
        v-show="totalPages!=0"
        v-bind:totalPages="totalPages"
        v-bind:currentPage="currentPage"
        v-bind:pageSize="pageSize"
        v-on:pagechanged="pageChangedClick"
        >
        </Pagination>
        <br>
        <span v-if="!users">No hay usuarios en este grupo</span>

    </div>

    <SideNav
    class="container-sidenav"
    v-bind:links="routesNav"
    >
    </SideNav>
    
</main>
</template>

<script>
// @ is an alias to /src
import SideNav from '@/components/SideNav.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'
import Pagination from '@/components/Pagination.vue'

import {formatter} from '@/components/mixin/formatter.js'
import {pagination} from '@/components/mixin/pagination.js'
import {restGroup} from '@/components/mixin/restGroup.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restGroup],
    data(){
        return{
            users: []
        }
    },
    components:{
        SideNav,
        Breadcrumb,
        Pagination
    },
    computed: {
        routesNav: function(){
          return [
                {name:'Categorías'},
                {name:'Editar grupo' , path: '/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup + '/editar-grupo'},
                {name:'Historial de cambios' , path: '/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup + '/historial-de-cambios'},
                {name:'Eliminar grupo' , path: '/usuarios/grupos-de-usuarios/' + this.$route.params.idGroup + '/eliminar-grupo'}
              ]
        },
        routesBread: function(){
          return [
                {name: 'Usuarios', path: '/usuarios'},
                {name: 'Grupos', path: '/usuarios/grupos-de-usuarios/'},
                {name: this.group.name},
          ]
        }
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            this.getGroupAllUsers(this.currentPage);
        },
        getGroupAllUsers: function(page){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/' + this.$route.params.idGroup + '/users?page=' + page) //Paginación
            .then(function (response) {
                that.users = response.data;
                response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
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
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroup();
        this.getGroupAllUsers(this.currentPage);
    }
}
</script>

<style lang="scss" scoped>

</style>
