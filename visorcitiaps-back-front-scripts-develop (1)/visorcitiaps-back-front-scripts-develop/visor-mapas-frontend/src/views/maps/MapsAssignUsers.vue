<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Asignar usuarios a {{map.name}}</h2>
    </div>

    <div class="container-action-center">
        
    </div>

    <div class="container-body">

        <h3 style="margin-bottom:0px">Todos los usuarios</h3>

        <div class="search-tools">
            <select id="group"
            v-model="selectedGroup"
            v-on:change="onChangeGroup(selectedGroup)">
                <option
                value="todas"
                >
                    Cualquier usuario
                </option>
                <option
                v-for="item in groups"
                v-bind:key="item.index"
                v-bind:value="item.id">
                {{item.name}}
                </option>
            </select>

            <input id="firstname" v-model="inputUser" type="text" placeholder="Buscar usuario..."/>
            <span>
                <button :disabled="!isAssignAllEnabled" class="btn btn-secondary" v-on:click="postManyRelMapLayer()">Asignar todas</button>
            </span>
        </div>

<!--         <div v-if="!showFilterResults" class="list-item-box">
            <template 
            v-for="(item) in allUsers" >
            <div class="list-item-item">
                <span v-bind:key="item.index" class="list-item-name">{{item.firstname + item.lastname}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.username}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button v-bind:key="item.index" v-if="!item.isAssigned" class="btn btn-secondary" v-on:click="postRelMapLayer(item)">Asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span></div>
            </template>
        </div>

        <div v-else class="list-item-box">
            <template 
            v-for="(item) in allLayersFiltered">
            <div class="list-item-item">
                <span v-bind:key="item.index" class="list-item-name">{{item.firstname + item.lastname}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.username}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button v-bind:key="item.index" v-if="!item.isAssigned" class="btn btn-secondary" v-on:click="postRelMapLayer(item)">Asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span></div>
            </template>
        </div> -->

        <assing-box  v-if="!showFilterResults" 
        :list="allUsers" 
        :title="['firstname']" 
        :title2="['lastname']"
        :subTitle="['username']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postRelMapLayer" 
        :assingText="'Asignar'">
        </assing-box>

        <assing-box v-else 
        :list="allLayersFiltered" 
        :title="['firstname']" 
        :title2="['lastname']"
        :subTitle="['username']" 
        :isAssingned="['isAssigned']" 
        @assingMethod="postRelMapLayer" 
        :assingText="'Asignar'">
        </assing-box>

    </div>

    <div class="container-list">

        <h3 style="margin-bottom:0px">Usuarios asignados</h3>

        <div class="search-tools">
            <span>
                <button :disabled="assignedUsers==null? true:false" class="btn btn-secondary" v-on:click="deleteManyRelMapLayer()">No asignar todas</button>
            </span>
            <div></div>
            <span>
                <button class="btn btn-secondary" v-on:click="finishAssignment()">Finalizar asignación</button>
            </span>
        </div>

<!--         <div class="list-item-box">
            <template 
            v-for="(item) in assignedUsers" >
            <div class="list-item-item">
                <span v-bind:key="item.index" class="list-item-name">{{item.firstname + item.lastname}}</span>
                <span v-bind:key="item.index" class="list-item-sub-1">{{item.username}}</span>
                <span v-bind:key="item.index" class="list-item-sub-3">
                    <button class="btn btn-secondary" v-on:click="deleteRelMapLayer(item)">No asignar</button>
                </span>
                <span v-bind:key="item.index" class="list-item-separator"></span></div>
            </template>

        </div> -->
        <assing-box
        :list="assignedUsers" 
        :title="['firstname']" 
        :title2="['lastname']"
        :subTitle="['username']" 
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
            groups: [],

            allUsers: [],
            assignedUsers: [],

            selectedGroup: 'todas',
            inputUser: '',
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
                if(this.allUsers == null) {
                    return []
                }
                list = this.allUsers;
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
            if (!this.allUsers) {
                return
            }
            return this.allUsers.filter(item => {
                return item.username.toLowerCase().includes(this.inputUser.toLowerCase()) ||
                item.firstname.toLowerCase().includes(this.inputUser.toLowerCase()) ||
                item.lastname.toLowerCase().includes(this.inputUser.toLowerCase())
            })
        }
    },
    watch: {
        'inputUser': function(){
            /*
                Si se ingresa algún texto que no sean solo espacios
            */
            if(/\S/.test(this.inputUser))
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
            this.getGroups(this.currentPage);
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
        getGroups: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/') //Sin paginación
            .then(function (response) {
                that.groups = response.data;
                that.getAssignedUsers();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAssignedUsers: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/users') //Sin paginación
            .then(function (response) {
                that.assignedUsers = response.data;
                that.onChangeGroup(that.selectedGroup);
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAllUsers: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/users/') //Sin paginación
            .then(function (response) {
                that.allUsers = response.data;
                that.checkAlreadyAssignedUsers();
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getGroupsAllUsers: function(idGroup){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/' + idGroup + '/users') //Sin paginación
            .then(function (response) {
                that.allUsers = response.data;
                //that.allUsers.forEach(function(item) { item.isAssigned = false; });
                that.checkAlreadyAssignedUsers();
            })
            .catch(function (error) {
                console.log(error)
                //that.message = error.response.data.message;
            });
        },
        checkAlreadyAssignedUsers: function(){
            if(this.allUsers == null || this.assignedUsers == null) {
                //Si alguna está vacía, no es necesario ejecutar esta función
                return;
            }
            var assignedLayersIndexInAllLayers = [];
            for (let i = 0; i < this.assignedUsers.length; i++) {
                let index = this.allUsers.map(function(item){
                     return item.id; 
                     }).indexOf(this.assignedUsers[i].id);
                assignedLayersIndexInAllLayers.push(index);
            }
            for (let i = 0; i < assignedLayersIndexInAllLayers.length; i++) {
                if(assignedLayersIndexInAllLayers[i] >= 0) {
                    this.allUsers[assignedLayersIndexInAllLayers[i]].isAssigned = true;
                }
            }
            //this.$forceUpdate(); //Actualiza la vista
        },
        postRelMapLayer: function(layer){
            var that = this;
            this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/users/' + layer.id, 
            {
                id_map: this.$route.params.idMap,
                id_user: layer.id
            })
            .then(function (response) {
                //Actualizar datos
                that.getAssignedUsers();
            })
            .catch(function (error) {
                console.log(error);
                that.message = error.response.data.message;
            });
        },
        deleteRelMapLayer: function(layer){
            var that = this;
            console.log(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/users/' + layer.id)
            this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/users/' + layer.id)
            .then(function (response) {
                //Actualizar datos
                that.getAssignedUsers();
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
                list = this.allUsers;
            }
            for (let i = 0; i < list.length; i++) {

                if(!list[i].isAssigned) {
                    this.$http.post(this.$store.getters.backendurl + 'api/v1/maps/' + this.map.id + '/users/' +  list[i].id, 
                {
                    id_map: this.map.id,
                    id_user: list[i].id
                })
                .then(function (response) {
                    that.getAssignedUsers();
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
            for (let i = 0; i < this.assignedUsers.length; i++) {

                this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.map.id + '/users/' + this.assignedUsers[i].id , 
                {
                    id_map: this.map.id,
                    id_user: this.assignedUsers[i].id
                })
                .then(function (response) {
                    that.getAssignedUsers();
                })
                .catch(function (error) {
                    console.log(error);
                    that.message = error.response.data.message;
                });
                
            }
        },
        onChangeGroup: function(idGroup){
            if(idGroup == 'todas') {
                this.getAllUsers();
            } else {
                this.getGroupsAllUsers(idGroup);
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
        this.getGroups();
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
