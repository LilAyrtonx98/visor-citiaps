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
        <h2>Grupos</h2>
    </div>

    <div class="container-body">

        <div class="list-item-area">
            <!--
            <template 
            v-for="(item, index) in layers" >
            -->
            <template 
            v-for="(item) in groups" >
                <span v-bind:key="item.index" class="list-item-name">
                    <router-link :to="'/usuarios/grupos-de-usuarios/' + item.id">{{item.name}}</router-link>
                </span>
                <span v-bind:key="item.index" class="list-item-sub-1" style="grid-column-end: 2;">{{dateFormatter(item.created_at)}}</span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
                
            </template>

            <span v-if="!groups">No hay grupos</span>

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
                {name:'Grupos'},
                {name:'Nuevo grupo' , path: '/usuarios/grupos-de-usuarios/nuevo-grupo'}
              ]
        },
        routesBread: function(){
          return [
                {name: 'Usuarios', path: '/usuarios'},
                {name: 'Grupos'},
          ]
        }
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            this.getGroupsPages(this.currentPage);
        },
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroupsPages(this.currentPage)
    }
}
</script>

<style lang="scss" scoped>

</style>
