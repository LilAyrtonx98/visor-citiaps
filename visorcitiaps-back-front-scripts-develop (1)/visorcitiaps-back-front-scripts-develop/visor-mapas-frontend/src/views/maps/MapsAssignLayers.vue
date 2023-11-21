<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Asignar capas a {{map.name}}</h2>
    </div>

    <div class="container-action-center">
        
    </div>

    <div class="container-body">

        <h3 style="margin-bottom:0px">Todas las capas</h3>

        <div class="search-tools">
            <select id="group"
            v-model="selectedCategory"
            v-on:change="onChangeGroup(selectedCategory)">
                <option
                value="todas"
                >
                    Cualquier capa
                </option>
                <option
                v-for="item in categories"
                v-bind:key="item.index"
                v-bind:value="item.id">
                {{item.name}}
                </option>
            </select>

            <input id="firstname" v-model="inputLayer" type="text" placeholder="Buscar capa..."/>
            <span>
                <button :disabled="!isAssignAllEnabled" class="btn btn-secondary" v-on:click="postManyRelMapLayer()">Asignar todas</button>
            </span>
        </div>

<!--         <div v-if="!showFilterResults" class="list-item-area">
            <template 
            v-for="(item) in allLayers" >
            <div class="list-item-item">
                <span v-bind:key="item.index" class="list-item-name">{{item["name"]}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item["geoserverdata"]["filename"]}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button v-bind:key="item.index" v-if="!item.isAssigned" class="btn btn-secondary" v-on:click="postRelMapLayer(item)">Asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span></div>
            </template>
        </div>

        <div v-else class="list-item-area">
            <template 
            v-for="(item) in allLayersFiltered">
                <span v-bind:key="item.index" class="list-item-name">{{item["name"]}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item["geoserverdata"]["filename"]}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button v-bind:key="item.index" v-if="!item.isAssigned" class="btn btn-secondary" v-on:click="postRelMapLayer(item)">Asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>
        </div> -->
        <assing-box  v-if="!showFilterResults" 
        :list="allLayers" 
        :title="['name']"
        :isAssingned="['isAssigned']" 
        @assingMethod="postRelMapLayer" 
        :assingText="'Asignar'">
        </assing-box>

        <assing-box v-else 
        :list="allLayersFiltered" 
        :title="['name']" 
        :subTitle="['geoserverdata']"
        :subTitle2="['filename']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postRelMapLayer" 
        :assingText="'Asignar'">
        </assing-box>
        

    </div>

    <div class="container-list">

        <h3 style="margin-bottom:0px">Capas asignadas</h3>

        <div class="search-tools">
            <span>
                <button :disabled="assignedLayers==null? true:false" class="btn btn-secondary" v-on:click="deleteManyRelMapLayer()">No asignar todas</button>
            </span>
            <div></div>
            <span>
                <button class="btn btn-secondary" v-on:click="finishAssignment()">Finalizar asignación</button>
            </span>
        </div>

<!--         <div class="list-item-area">
            <template 
            v-for="(item) in assignedLayers" >
                <span v-bind:key="item.index" class="list-item-name">{{item.name}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.geoserverdata.filename}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button class="btn btn-secondary" v-on:click="deleteRelMapLayer(item)">No asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span>
            </template>

        </div> -->
        <assing-box
        :list="assignedLayers" 
        :title="['name']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="deleteRelMapLayer" 
        :assingText="'No asignar'">
        </assing-box>
        
    </div>
    
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'
import assingBox from '@/components/AssingBox.vue'

export default{
    name: 'test',
    data(){
        return{
            map: '',
            categories: [],

            allLayers: [],
            assignedLayers: [],

            selectedCategory: 'todas',
            inputLayer: '',
            showFilterResults: false,

            currentPage: 1,
            totalPages: 0,
            pageSize: 15,
        }
    },
    components:{
        Breadcrumb,
        assingBox
    },
    computed: {
        isAssignAllEnabled: function() {

            var list;
            let count = 0;
            if(this.showFilterResults) {
                list = this.allLayersFiltered;
            } else {
                if(this.allLayers == null) {
                    return []
                }
                list = this.allLayers;
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
              {name: this.map.name, path: '/mapas/' + this.$route.params.idMap},
              {name: 'Asignar capas'}
              ]
        },
        allLayersFiltered: function () {
            if (!this.allLayers) {
                return
            }
            return this.allLayers.filter(item => {
                return item.name.toLowerCase().includes(this.inputLayer.toLowerCase()) ||
                item.geoserverdata.filename.toLowerCase().includes(this.inputLayer.toLowerCase())
            })
        }
    },
    watch: {
        'inputLayer': function(){
            /*
                Si se ingresa algún texto que no sean solo espacios
            */
            if(/\S/.test(this.inputLayer))
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
            this.getCategories(this.currentPage);
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
        getCategories: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/') //Sin paginación
            .then(function (response) {
                that.categories = response.data;
                that.getAssignedLayers();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAssignedLayers: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/layers') //Sin paginación
            .then(function (response) {
                that.assignedLayers = response.data;
                that.onChangeGroup(that.selectedCategory);
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAllLayers: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/layers/') //Sin paginación
            .then(function (response) {
                that.allLayers = response.data;
                //that.allLayers.forEach(function(item) { item.isAssigned = false; });
                that.checkAlreadyAssignedLayers();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getCategoryAllLayers: function(idCategory){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/' + idCategory + '/layers') //Sin paginación
            .then(function (response) {
                that.allLayers = response.data;
                //that.allLayers.forEach(function(item) { item.isAssigned = false; });
                that.checkAlreadyAssignedLayers();
            })
            .catch(function (error) {
                console.log(error)
                //that.message = error.response.data.message;
            });
        },
        checkAlreadyAssignedLayers: function(){
            if(this.allLayers == null || this.assignedLayers == null) {
                //Si alguna está vacía, no es necesario ejecutar esta función
                return;
            }
            var assignedLayersIndexInAllLayers = [];
            for (let i = 0; i < this.assignedLayers.length; i++) {
                let index = this.allLayers.map(function(item){
                     return item.id; 
                     }).indexOf(this.assignedLayers[i].id);
                assignedLayersIndexInAllLayers.push(index);
            }
            for (let i = 0; i < assignedLayersIndexInAllLayers.length; i++) {
                if(assignedLayersIndexInAllLayers[i] >= 0) {
                    this.allLayers[assignedLayersIndexInAllLayers[i]].isAssigned = true;
                }
            }
            //this.$forceUpdate(); //Actualiza la vista
        },
        postRelMapLayer: function(layer){
            var that = this;
            this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/layers/' + layer.id, 
            {
                id_map: this.$route.params.idMap,
                id_layer: layer.id
            })
            .then(function (response) {
                //Actualizar datos
                that.getAssignedLayers();
            })
            .catch(function (error) {
                console.log(error);
                that.message = error.response.data.message;
            });
        },
        deleteRelMapLayer: function(layer){
            var that = this;
            this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/layers/' + layer.id)
            .then(function (response) {
                //Actualizar datos
                that.getAssignedLayers();
            })
            .catch(function (error) {
                console.log(error);
                that.message = error.response.data.message;
            });
        },
        postManyRelMapLayer: function(){
            var that = this;
            var list;
            if(this.showFilterResults) {
                list = this.allLayersFiltered;
            } else {
                list = this.allLayers;
            }
            for (let i = 0; i < list.length; i++) {

                if(!list[i].isAssigned) {
                    this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + this.map.id + '/layers/' +  list[i].id, 
                {
                    id_map: this.map.id,
                    id_layer: list[i].id
                })
                .then(function (response) {
                    that.getAssignedLayers();
                })
                .catch(function (error) {
                    console.log(error);
                    that.message = error.response.data.message;
                });
                }
            }
        },
        deleteManyRelMapLayer: function(){
            var that = this;
            for (let i = 0; i < this.assignedLayers.length; i++) {

                this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.map.id + '/layers/' + this.assignedLayers[i].id , 
                {
                    id_map: this.map.id,
                    id_layer: this.assignedLayers[i].id
                })
                .then(function (response) {
                    that.getAssignedLayers();
                })
                .catch(function (error) {
                    console.log(error);
                    that.message = error.response.data.message;
                });
                
            }
        },
        onChangeGroup: function(idCategory){
            if(idCategory == 'todas') {
                this.getAllLayers();
            } else {
                this.getCategoryAllLayers(idCategory);
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
        this.getCategories();
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
