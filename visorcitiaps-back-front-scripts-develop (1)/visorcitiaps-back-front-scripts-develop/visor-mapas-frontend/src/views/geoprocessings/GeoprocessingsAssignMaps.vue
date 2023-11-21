<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Asignar mapas a {{geo.name}}</h2>
    </div>

    <div class="container-action-center">
        
    </div>

    <div class="container-body">

        <h3 style="margin-bottom:0px">Todos los mapas</h3>

        <div class="search-tools" style="grid-template-columns: 1fr max-content;">

            <input id="firstname" v-model="inputMap" type="text" placeholder="Buscar geoproceso..."/>
            <span>
                <button :disabled="!isAssignAllEnabled" class="btn btn-secondary" v-on:click="postManyRelMapGeoprocessing()">Asignar todos</button>
            </span>
        </div>

<!--         <div v-if="!showFilterResults" class="list-item-area">
            <template 
            v-for="(item) in allMaps" >
                <span v-bind:key="item.index" class="list-item-name">{{item.name}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.geo_url}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button v-bind:key="item.index" v-if="!item.isAssigned" class="btn btn-secondary" v-on:click="postRelMapGeoprocessing(item)">Asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>
        </div>

        <div v-else class="list-item-area">
            <template 
            v-for="(item) in allMapsFiltered">
                <span v-bind:key="item.index" class="list-item-name">{{item.name}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.geo_url}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button v-bind:key="item.index" v-if="!item.isAssigned" class="btn btn-secondary" v-on:click="postRelMapGeoprocessing(item)">Asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>
        </div> -->

        <assing-box  v-if="!showFilterResults" 
        :list="allMaps" 
        :title="['name']" 
        :subTitle="['geo_url']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postRelMapGeoprocessing" 
        :assingText="'Asignar'">
        </assing-box>

        <assing-box v-else 
        :list="allMapsFiltered" 
        :title="['name']" 
        :subTitle="['geo_url']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postRelMapGeoprocessing" 
        :assingText="'Asignar'">
        </assing-box>

    </div>

    <div class="container-list">

        <h3 style="margin-bottom:0px">Mapas asignados</h3>

        <div class="search-tools">
            <span>
                <button :disabled="assignedMaps==null? true:false" class="btn btn-secondary" v-on:click="deleteManyRelMapGeoprocessing()">No asignar todas</button>
            </span>
            <div></div>
            <span>
                <button class="btn btn-secondary" v-on:click="finishAssignment()">Finalizar asignación</button>
            </span>
        </div>

<!--         <div class="list-item-area">
            <template 
            v-for="(item) in assignedMaps" >
                <span v-bind:key="item.index" class="list-item-name">{{item.name}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.geo_url}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button class="btn btn-secondary" v-on:click="deleteRelMapGeoprocessing(item)">No asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>

        </div> -->

        <assing-box
        :list="assignedMaps" 
        :title="['name']" 
        :subTitle="['geo_url']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="deleteRelMapGeoprocessing" 
        :assingText="'No asignar'">
        </assing-box>

        <Pagination
        v-show="totalPages!=0"
        v-bind:totalPages="totalPages"
        v-bind:currentPage="currentPage"
        v-bind:pageSize="pageSize"
        v-on:pagechanged="pageChangedClick"
        >
        </Pagination>
        
    </div>
    
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'
import Pagination from '@/components/Pagination.vue'
import assingBox from '@/components/AssingBox.vue'

export default{
    name: 'test',
    data(){
        return{
            geo: '',

            allMaps: [],
            assignedMaps: [],

            inputMap: '',
            showFilterResults: false,

            currentPage: 1,
            totalPages: 0,
            pageSize: 15,
        }
    },
    components:{
        Breadcrumb,
        Pagination,
        assingBox
    },
    computed: {
        isAssignAllEnabled: function() {

            var list;
            let count = 0;
            if(this.showFilterResults) {
                list = this.allMapsFiltered;
            } else {
                if(this.allMaps == null) {
                    return []
                }
                list = this.allMaps;
            }
            for (let i = 0; i < list.length; i++) {
                if(!list[i].isAssigned) {
                    count++;
                }
            }
            if(count>0) {
                return true;
            }
            return false;
        },
        routesBread: function(){
          return [
              {name:'Mapas', path: '/mapas' },
              {name: this.geo.name, path: '/mapas/' + this.$route.params.idGeo},
              {name: 'Asignar mapas'}
              ]
        },
        allMapsFiltered: function () {
            if (!this.allMaps) {
                return
            }
            return this.allMaps.filter(item => {
                return item.name.toLowerCase().includes(this.inputMap.toLowerCase())
            })
        }
    },
    watch: {
        'inputMap': function(){
            /*
                Si se ingresa algún texto que no sean solo espacios
            */
            if(/\S/.test(this.inputMap))
            {
                this.showFilterResults = true;
            }
            else
            {
                this.showFilterResults = false;
            }
        },
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
        },
        getGeoprocessing: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/geoprocessings/' + this.$route.params.idGeo)
            .then(function (response) {
                that.geo = response.data;
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAssignedMaps: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/geoprocessings/' + this.$route.params.idGeo + '/maps') //Sin paginación
            .then(function (response) {
                that.assignedMaps = response.data;
                that.getAllMaps();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAllMaps: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/maps') //Sin paginación
            .then(function (response) {
                that.allMaps = response.data;
                //that.allMaps.forEach(function(item) { item.isAssigned = false; });
                that.checkAlreadyassignedMaps();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        checkAlreadyassignedMaps: function(){
            if(this.allMaps == null || this.assignedMaps == null) {
                //Si alguna está vacía, no es necesario ejecutar esta función
                return;
            }
            var assignedGeoprocessingsIndexInallGeoprocessings = [];
            for (let i = 0; i < this.assignedMaps.length; i++) {
                let index = this.allMaps.map(function(item){
                     return item.id; 
                     }).indexOf(this.assignedMaps[i].id);
                assignedGeoprocessingsIndexInallGeoprocessings.push(index);
            }
            for (let i = 0; i < assignedGeoprocessingsIndexInallGeoprocessings.length; i++) {
                if(assignedGeoprocessingsIndexInallGeoprocessings[i] >= 0) {
                    this.allMaps[assignedGeoprocessingsIndexInallGeoprocessings[i]].isAssigned = true;
                }
            }
            //this.$forceUpdate(); //Actualiza la vista
        },
        postRelMapGeoprocessing: function(map){
            var that = this;
            this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + map.id + '/geoprocessings/' + this.$route.params.idGeo, 
            {
                id_map: map.id,
                id_geoprocessing: this.$route.params.idGeo
            })
            .then(function (response) {
                //Actualizar datos
                that.getAssignedMaps();
            })
            .catch(function (error) {
                console.log(error);
                that.message = error.response.data.message;
            });
        },
        deleteRelMapGeoprocessing: function(map){
            var that = this;
            this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + map.id + '/geoprocessings/' + this.$route.params.idGeo)
            .then(function (response) {
                //Actualizar datos
                that.getAssignedMaps();
            })
            .catch(function (error) {
                console.log(error);
                that.message = error.response.data.message;
            });
        },
        postManyRelMapGeoprocessing: function(){
            var that = this;
            var list;
            if(this.showFilterResults) {
                list = this.allMapsFiltered;
            } else {
                list = this.allMaps;
            }
            for (let i = 0; i < list.length; i++) {

                if(!list[i].isAssigned) {
                    this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + list[i].id + '/geoprocessings/' +  this.geo.id, 
                {
                    id_map: list[i].id,
                    id_geoprocessing: this.geo.id
                })
                .then(function (response) {
                    that.getAssignedMaps();
                })
                .catch(function (error) {
                    console.log(error);
                    that.message = error.response.data.message;
                });
                }
            }
        },
        deleteManyRelMapGeoprocessing: function(){
            var that = this;
            for (let i = 0; i < this.assignedMaps.length; i++) {

                this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.assignedMaps[i].id + '/geoprocessings/' +  this.geo.id, 
                {
                    id_map: this.assignedMaps[i].id,
                    id_geoprocessing: this.geo.id
                })
                .then(function (response) {
                    that.getAssignedMaps();
                })
                .catch(function (error) {
                    console.log(error);
                    that.message = error.response.data.message;
                });
                
            }
        },
        finishAssignment: function(){
            this.$router.push('/geoprocesos/' + this.$route.params.idGeo)
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getGeoprocessing();
        this.getAssignedMaps();
        this.getAllMaps();
    }
}
</script>

<style lang="scss" scoped>
.search-tools{
    display: grid;
    grid-template-columns: 1fr 1fr max-content;
    grid-column-gap: 0.5em;
}
input, select{
    width: 0.5fr;
}
@media only screen and (max-width: 640px) {
    .search-tools{
        grid-template-columns: 1fr;
        grid-row-gap: 0.5em;
    }
}
</style>
