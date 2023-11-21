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
        <h2>Categorías</h2>
    </div>

    <div class="container-body">

        <div class="list-item-area">
            <template 
            v-for="(item) in categories" >
                <span v-bind:key="item.index" class="list-item-name">
                    <router-link :to="'/capas/categorias-de-capas/' + item.id">{{item.name}}</router-link>
                </span>
                <span v-bind:key="item.index" class="list-item-sub-1" style="grid-column-end: 2;">{{dateFormatter(item.created_at)}}</span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>

            <span v-if="!categories">No hay categorías</span>

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
import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restCategory],
    data(){
        return{
            categories: []
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
                {name:'Nueva categoría' , path: '/capas/categorias-de-capas/nueva-categoria'},
              ]
        },
        routesBread: function(){
          return [
              {name: 'Capas', path: '/capas' },
              {name: 'Categorías'}
          ]
        }
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            this.getCategoriesPages(this.currentPage);
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    }, 
    mounted () {
        this.getCategoriesPages(this.currentPage);
    }
}
</script>

<style lang="scss" scoped>

</style>
