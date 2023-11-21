<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Lista de capas</h2>
        
    </div>
    
    <div class="container-body">

        <div class="list-item-area">
            <div class="search-tools">
                <select id="category"
                v-model="selectedCategory"
                v-on:change="onChangeCategory(selectedCategory)">
                <option value="todas">
                    Cualquier capa
                </option>
                <option v-for="item in categories"
                v-bind:key="item.index"
                v-bind:value="item.id">
                {{item.name}}
                </option>
                </select>
            </div>
            <template 
            v-for="(item) in layers" >
                <span v-bind:key="item.index" class="list-item-name">
                    <router-link :to="'/capas/' + item.id"  >{{item.name}}</router-link>
                </span>
                <span v-if="item.geoserverdata" v-bind:key="item.index" class="list-item-sub-1" style="grid-column-end: 3;">{{item.geoserverdata.filename}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">{{layerCategoryName(item)}}</span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>
            <span v-if="!layers">No hay capas</span>
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
import {restLayer} from '@/components/mixin/restLayer.js'
import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restLayer, restCategory],
    data(){
        return{
            message: '',
            selectedCategory: 'todas'
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
                {name:'Capas'},
                {name:'Nueva capa' , path: '/capas/nueva-capa'}, 
                {name:'Categorías'},
                {name:'Ver categorías' , path: '/capas/categorias-de-capas'},
              ]
        },
        routesBread: function(){
          return [{name:'Capas' }]
        },
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            if(this.selectedCategory=="todas"){
                this.getLayersPages(this.currentPage);
            }
            else{
                this.getLayersByCategory(this.currentPage, this.selectedCategory)
            }
        },
        layerCategoryName: function(layer){
            var category = this.categories.find(function(element) {
                    return element.id == layer.id_category;
                });
            if(category == undefined)
            {
                return 'Categoría desconocida';
            }
            return category.name;
        },
        onChangeCategory: function(idCaregory){
            this.currentPage=1
            if(idCaregory == 'todas'){
                this.getLayersPages(this.currentPage)
            }
            else{
                this.getLayersByCategory(this.currentPage,idCaregory)
            }
        },
        test: function(){
            console.log('TEST')
        }
    },
    updated () {

    },
    created () {
        //Verificar permisos
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    beforeMount() {
        
    },
    mounted () {
        this.getCategoriesNoPages();
        this.getLayersPages(this.currentPage);
    }
}
</script>

<style lang="scss" scoped>

</style>
