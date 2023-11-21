<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{layer.name}}</h2>
    </div>

    <div class="container-body">
        <strong>Información</strong>
        <p>{{category.name}}</p>
        <p>{{dateFormatter(layer.created_at)}}</p>
        <p v-if="layer.desc">{{layer.desc}}</p>
        <br>
        <strong>Datos GeoServer</strong>
        <p> <label for="filename">Nombre de archivo: </label> <span id="filename">{{layer.provider.geoserverdata.filename}}</span></p>
        
        <p><label for="coordsystem">Sistema de coordenadas: </label> <span id="coordsystem">{{layer.provider.geoserverdata.coordinatessystem}}</span></p>

        <p><label for="workspace">Espacio de trabajo: </label> <span id="workspace">{{layer.provider.geoserverdata.workspace}}</span></p>
        
        <p><label for="datastore">Almacen de datos: </label> <span id="datastore">{{layer.provider.geoserverdata.datastore}}</span></p>

        <br>
        <strong>Mapas en los cuales esta capa se encuentra asignada</strong>
        <div class="list-item-area">
            <template 
            v-for="(item) in calculateAssignedMapsPage" >
                <router-link 
                v-bind:key="item.index" 
                class="list-item-name"
                :to="'/mapas/' + item.id"
                >
                    {{item.name}}
                </router-link>
                <span v-bind:key="item.index" class="list-item-sub-1"></span>
                <span v-bind:key="item.index" class="list-item-sub-3"></span>
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
        <span v-if="!assignedMaps">No hay usuarios asignados</span>

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
import {restGroup} from '@/components/mixin/restGroup.js'
import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restLayer, restGroup, restCategory],
    data(){
        return{
            assignedMaps: [],
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
                {name:'Editar capa' , path: '/capas/' + this.$route.params.idLayer + '/editar-capa'},
                {name:'Asignar mapas' , path: '/capas/' + this.$route.params.idLayer + '/asignar-mapas'},
                {name:'Historial de cambios' , path: '/capas/' + this.$route.params.idLayer + '/historial-de-cambios'},
                {name:'Eliminar capa' , path: '/capas/' + this.$route.params.idLayer + '/eliminar-capa'},
              ]
        },
        routesBread: function(){
          return [
              {name:'Capas', path: '/capas' },
              {name: this.layer.name}
              ]
        },
        calculateAssignedMapsPage: function (){

            if(this.assignedMaps == null) {
                //Si alguna está vacía, no es necesario ejecutar esta función
                return;
            }
            let pageResults = [];

            let skip = this.currentPage > 0 ? ( ( this.currentPage - 1 ) * this.pageSize ) : 0 ;
            for (let i = skip; i < this.assignedMaps.length && i < skip + this.pageSize ; i++) {
                pageResults.push(this.assignedMaps[i])
            }

            //Calcula total de páginas
            if(this.assignedMaps.length <= this.pageSize) {
                this.totalPages = 0;
                return pageResults;
            }
            if ( this.assignedMaps.length % this.pageSize > 0 ) {
                this.totalPages =  Math.floor( (this.assignedMaps.length / this.pageSize) + 1 )
                return pageResults;
            } else {
                this.totalPages = Math.floor( (this.assignedMaps.length / this.pageSize) )
                return pageResults;
            }
        },
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
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
        getCategoryByID: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/' + this.layer.id_category)
            .then(function (response) {
                that.category = response.data;
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAssignedUsers: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer + '/maps') //Sin paginación
            .then(function (response) {
                that.assignedMaps = response.data;
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
        this.getGroupsNoPages();
        this.getLayer();
        this.getAssignedUsers();
    },
    watch: {
        'layer': function (newValue, oldValue) {
            if(newValue.id_category){
                this.getCategoryByID();
            }
        },
    }
}
</script>

<style lang="scss" scoped>

</style>
