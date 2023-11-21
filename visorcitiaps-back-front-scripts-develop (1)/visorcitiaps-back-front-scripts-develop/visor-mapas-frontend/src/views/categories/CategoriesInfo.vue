<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{category.name}}</h2>
    </div>

    <div class="container-body">
        <strong>Información</strong>
        <p>{{dateFormatter(category.created_at)}}</p>
        <p>{{category.desc}}</p>
        <br>
        <strong>Capas que pertenecen a esta categoría</strong>
        <div class="list-item-area">
            <template 
            v-for="(item) in layers" >
                <router-link 
                v-bind:key="item.index" 
                class="list-item-name"
                :to="'/capas/' + item.id"
                >
                    {{item.name}}
                </router-link>
                <span v-if="item.geoserverdata" v-bind:key="item.index" class="list-item-sub-1" style="grid-column-end: 3;">{{item.geoserverdata.filename}}</span>
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
        <span v-if="!layers">No hay capas en esta categoría</span>

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

            layers: []
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
                {name:'Editar categoría' , path: '/capas/categorias-de-capas/' + this.$route.params.idCategory + '/editar-categoria'},
                {name:'Historial de cambios' , path: '/capas/categorias-de-capas/' + this.$route.params.idCategory + '/historial-de-cambios'},
                {name:'Eliminar categoría' , path: '/capas/categorias-de-capas/' + this.$route.params.idCategory + '/eliminar-categoria'},
              ]
        },
        routesBread: function(){
          return [
                {name: 'Capas', path: '/capas'},
                {name: 'Categorías', path: '/capas/categorias-de-capas/'},
                {name: this.category.name},
          ]
        }
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            this.getCategoryAllLayers(this.currentPage);
        },
        getCategoryAllLayers: function(page){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/' + this.$route.params.idCategory + '/layers?page=' + page) //Paginación
            .then(function (response) {
                that.layers = response.data;
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
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getCategory();
        this.getCategoryAllLayers(this.currentPage);
    }
}
</script>

<style lang="scss" scoped>

</style>
