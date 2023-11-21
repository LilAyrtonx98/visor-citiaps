<template>
<main class="container-right-nav">
    <!--
    <h2>Capas</h2>
    <img alt="Vue logo" src="../../assets/img/logo.png">
    -->

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Lista de usuarios</h2>
    </div>

    <div class="container-body">

        <div class="list-item-area">

        <div class="search-tools">
            <select id="group"
            v-model="selectedGroup"
            v-on:change="onChangeGroup(selectedGroup)">
            <option value="todas">
                Cualquier usuario
            </option>
            <option v-for="item in groups"
            v-bind:key="item.index"
            v-bind:value="item.id">
            {{item.name}}
            </option>
            </select>
        </div>
            <template 
            v-for="(item) in users" >
                <span v-bind:key="item.index" class="list-item-name">
                    <router-link :to="'/usuarios/' + item.id">{{item.firstname + ' ' + item.lastname}}</router-link>
                </span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.username}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">{{userGroupName(item)}}</span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>
            <span v-if="!users">No hay usuarios</span>
        </div>

        <Pagination
        v-show="totalPages!=0"
        v-bind:totalPages="totalPages"
        v-bind:currentPage="currentPage"
        v-bind:pageSize="pageSize"
        v-on:pagechanged="pageChangedClick"
        >
        </Pagination>

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
import Pagination from '@/components/Pagination.vue'

import {formatter} from '@/components/mixin/formatter.js'
import {pagination} from '@/components/mixin/pagination.js'
import {restUser} from '@/components/mixin/restUser.js'
import {restGroup} from '@/components/mixin/restGroup.js'

export default{
    mixins: [formatter, pagination, restUser, restGroup],
    data(){
        return{
            selectedGroup: 'todas'
            
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
                {name:'Usuarios'},
                {name:'Nuevo usuario' , path: '/usuarios/nuevo-usuario'}, 
                {name:'Grupos'},
                {name:'Ver grupos' , path: '/usuarios/grupos-de-usuarios'},
              ]
        },
        routesBread: function(){
          return [{name:'Usuarios' }]
        }
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            if(this.selectedGroup=="todas"){
                this.getUsersPages(this.currentPage);
            }
            else{
                this.getUsersByGroup(this.currentPage, this.selectedGroup)
            }
        },
        userGroupName: function(user){
            var group = this.groups.find(function(element) {
                    return element.id == user.id_group;
                });
            if(group == undefined)
            {
                return 'Grupo desconocido';
            }
            return group.name;
        },
        onChangeGroup: function(idGroup){
            this.currentPage=1
            if(idGroup == 'todas'){
                this.getUsersPages(this.currentPage)
            }
            else{
                this.getUsersByGroup(this.currentPage, idGroup)
            }
        }
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroupsNoPages();
        this.getUsersPages(this.currentPage);
    }
}
</script>
