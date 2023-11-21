<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Asignar geoprocesos a {{map.name}}</h2>
    </div>

    <div class="container-action-center">
        
    </div>

    <div class="container-body">

        <h3 style="margin-bottom:0px">Todos los geoprocesos</h3>

        <div class="search-tools" style="grid-template-columns: 1fr max-content;">

            <input id="firstname" v-model="inputGeoprocessing" type="text" placeholder="Buscar geoproceso..."/>
            <span>
                <button :disabled="!isAssignAllEnabled" class="btn btn-secondary" v-on:click="postManyRelMapGeoprocessing()">Asignar todos</button>
            </span>
        </div>
        <assing-box  v-if="!showFilterResults" 
        :list="allGeoprocessings" 
        :title="['name']" 
        :subTitle="['geo_url']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postManyRelMapGeoprocessing" 
        :assingText="'Asignar'">
        </assing-box>

        <assing-box v-else 
        :list="allUsersFiltered" 
        :title="['name']" 
        :subTitle="['geo_url']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postManyRelMapGeoprocessing" 
        :assingText="'Asignar'">
        </assing-box>

    </div>

    <div class="container-list">

        <h3 style="margin-bottom:0px">Geoprocesos asignados</h3>

        <div class="search-tools">
            <span>
                <button :disabled="assignedGeoprocessings==null? true:false" class="btn btn-secondary" v-on:click="deleteManyRelMapGeoprocessing()">No asignar todas</button>
            </span>
            <div></div>
            <span>
                <button class="btn btn-secondary" v-on:click="finishAssignment()">Finalizar asignación</button>
            </span>
        </div>

<!--         <div class="list-item-area">
            <template 
            v-for="(item) in assignedGeoprocessings" >
                <span v-bind:key="item.index" class="list-item-name">{{item.name}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.geo_url}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button class="btn btn-secondary" v-on:click="deleteRelMapGeoprocessing(item)">No asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>

        </div> -->
        <assing-box
        :list="assignedGeoprocessings" 
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
            map: '',

            allGeoprocessings: [],
            assignedGeoprocessings: [],

            inputGeoprocessing: '',
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
                list = this.allUsersFiltered;
            } else {
                if(this.allGeoprocessings == null) {
                    return []
                }
                list = this.allGeoprocessings;
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
        routesNav: function(){
          return [
                {name:'Categorías'},
                {name:'Nueva categoría' , path: '/capas/categorias/nueva-categoria'},
              ]
        },
        routesBread: function(){
          return [
              {name:'Mapas', path: '/mapas' },
              {name: this.map.name, path: '/mapas/' + this.$route.params.idMap},
              {name: 'Asignar geoprocesos'}
              ]
        },
        allUsersFiltered: function () {
            if (!this.allGeoprocessings) {
                return
            }
            return this.allGeoprocessings.filter(item => {
                return item.name.toLowerCase().includes(this.inputGeoprocessing.toLowerCase())
            })
        }
    },
    watch: {
        'inputGeoprocessing': function(){
            /*
                Si se ingresa algún texto que no sean solo espacios
            */
            if(/\S/.test(this.inputGeoprocessing))
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
        userGroupName: function(map){
            var group = this.groups.find(function(element) {
                    return element.id == map.id_group;
                });
            if(group == undefined)
            {
                return 'Grupo desconocido';
            }
            return group.name;
        },
        getMap: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap)
            .then(function (response) {
                that.map = response.data;
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getassignedGeoprocessings: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/geoprocessings') //Sin paginación
            .then(function (response) {
                that.assignedGeoprocessings = response.data;
                that.getAllGeoprocessings();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAllGeoprocessings: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/geoprocessings') //Sin paginación
            .then(function (response) {
                that.allGeoprocessings = response.data;
                //that.allGeoprocessings.forEach(function(item) { item.isAssigned = false; });
                that.checkAlreadyassignedGeoprocessings();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        checkAlreadyassignedGeoprocessings: function(){
            if(this.allGeoprocessings == null || this.assignedGeoprocessings == null) {
                //Si alguna está vacía, no es necesario ejecutar esta función
                return;
            }
            var assignedGeoprocessingsIndexInallGeoprocessings = [];
            for (let i = 0; i < this.assignedGeoprocessings.length; i++) {
                let index = this.allGeoprocessings.map(function(item){
                     return item.id; 
                     }).indexOf(this.assignedGeoprocessings[i].id);
                assignedGeoprocessingsIndexInallGeoprocessings.push(index);
            }
            for (let i = 0; i < assignedGeoprocessingsIndexInallGeoprocessings.length; i++) {
                if(assignedGeoprocessingsIndexInallGeoprocessings[i] >= 0) {
                    this.allGeoprocessings[assignedGeoprocessingsIndexInallGeoprocessings[i]].isAssigned = true;
                }
            }
            //this.$forceUpdate(); //Actualiza la vista
        },
        postRelMapGeoprocessing: function(geoprocessing){
            var that = this;
            this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/geoprocessings/' + geoprocessing.id, 
            {
                id_map: this.$route.params.idMap,
                id_geoprocessing: geoprocessing.id
            })
            .then(function (response) {
                //Actualizar datos
                that.getassignedGeoprocessings();
            })
            .catch(function (error) {
                console.log(error);
                that.message = error.response.data.message;
            });
        },
        deleteRelMapGeoprocessing: function(geoprocessing){
            var that = this;
            this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/geoprocessings/' + geoprocessing.id)
            .then(function (response) {
                //Actualizar datos
                that.getassignedGeoprocessings();
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
                list = this.allUsersFiltered;
            } else {
                list = this.allGeoprocessings;
            }
            for (let i = 0; i < list.length; i++) {

                if(!list[i].isAssigned) {
                    this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + this.map.id + '/geoprocessings/' +  list[i].id, 
                {
                    id_map: this.map.id,
                    id_geoprocessing: list[i].id
                })
                .then(function (response) {
                    that.getassignedGeoprocessings();
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
            for (let i = 0; i < this.assignedGeoprocessings.length; i++) {

                this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.map.id + '/geoprocessings/' + this.assignedGeoprocessings[i].id , 
                {
                    id_map: this.map.id,
                    id_geoprocessing: this.assignedGeoprocessings[i].id
                })
                .then(function (response) {
                    that.getassignedGeoprocessings();
                })
                .catch(function (error) {
                    console.log(error);
                    that.message = error.response.data.message;
                });
                
            }
        },
        finishAssignment: function(){
            this.$router.push('/mapas/' + this.$route.params.idMap)
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.maps? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getMap();
        this.getassignedGeoprocessings();
        this.getAllGeoprocessings();
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
