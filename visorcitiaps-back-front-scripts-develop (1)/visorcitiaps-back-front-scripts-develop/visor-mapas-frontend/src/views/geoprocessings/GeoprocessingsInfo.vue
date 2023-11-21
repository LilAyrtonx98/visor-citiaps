<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{geo.name}}</h2>
    </div>

    <div class="container-body">
        <strong>Información</strong>
        <p>{{geo.geo_url}}</p>
        <p>{{dateFormatter(geo.created_at)}}</p>
        <p v-if="geo.desc">{{geo.desc}}</p>
        <br>
        <strong>Mapas en los cuales este geoproceso se encuentra asignado</strong>
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
import {restGeoprocessing} from '@/components/mixin/restGeoprocessing.js'
import {restGroup} from '@/components/mixin/restGroup.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restGeoprocessing, restGroup],
    data(){
        return{
            assignedMaps: [],
            groups: [],
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
                {name:'Geoprocesos'},
                {name:'Editar geoproceso' , path: '/geoprocesos/' + this.$route.params.idGeo + '/editar-geoproceso'},
                {name:'Asignar mapas' , path: '/geoprocesos/' + this.$route.params.idGeo + '/asignar-mapas'},
                {name:'Historial de cambios' , path: '/geoprocesos/' + this.$route.params.idGeo + '/historial-de-cambios'},
                {name:'Eliminar geoproceso' , path: '/geoprocesos/' + this.$route.params.idGeo + '/eliminar-geoproceso'}
              ]
        },
        routesBread: function(){
          return [
              {name:'Geoprocesos', path: '/geoprocesos' },
              {name: this.geo.name},
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
        getAssignedMaps: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/geoprocessings/' + this.$route.params.idGeo + '/maps') //Sin paginación
            .then(function (response) {
                that.assignedMaps = response.data;
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.geo? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGroupsNoPages();
        this.getGeoprocessing();
        this.getAssignedMaps();
    }
}
</script>

<style lang="scss" scoped>

</style>
