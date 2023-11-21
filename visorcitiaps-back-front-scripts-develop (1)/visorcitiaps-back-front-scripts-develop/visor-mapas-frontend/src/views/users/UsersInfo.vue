<template>
<main class="container-right-nav">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>{{user.firstname + ' ' + user.lastname}}</h2>
    </div>

    <div class="container-body">
        <strong>Información</strong>
        <p>{{user.username}}</p>
        <p>{{userGroup.name}}</p>
        <p>{{dateFormatter(user.created_at)}}</p>
        <br>
        <strong>Permisos</strong>
        <div class="permissions">
            <p> 
                <font-awesome-icon v-if="user.permissions.users" icon="check"></font-awesome-icon>
                <font-awesome-icon v-else icon="times"></font-awesome-icon> 
                Gestión de usuarios
            </p>
            <p>
                <font-awesome-icon v-if="user.permissions.layers" icon="check"></font-awesome-icon>
                <font-awesome-icon v-else icon="times"></font-awesome-icon> 
                Gestión de capas
            </p>
            <p>
                <font-awesome-icon v-if="user.permissions.geo" icon="check"></font-awesome-icon>
                <font-awesome-icon v-else icon="times"></font-awesome-icon> 
                Gestión de geoprocesos
            </p>
            <p>
                <font-awesome-icon v-if="user.permissions.maps" icon="check"></font-awesome-icon>
                <font-awesome-icon v-else icon="times"></font-awesome-icon> 
                Gestión de mapas
            </p>
            <p>
                <font-awesome-icon v-if="user.permissions.visor" icon="check"></font-awesome-icon>
                <font-awesome-icon v-else icon="times"></font-awesome-icon> 
                Acceso al visor
            </p>
        </div>
        <br>
        <strong>Mapas asignados al usuario</strong>

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
        <span v-if="!logs">No hay acciones que mostrar</span>
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
import LogItem from '@/components/LogItem.vue'

import {formatter} from '@/components/mixin/formatter.js'
import {pagination} from '@/components/mixin/pagination.js'
import {restUser} from '@/components/mixin/restUser.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restUser],
    data(){
        return{
            userGroup: '',
            logs: [],
            groups: [],
            assignedMaps: []
        }
    },
    components:{
        SideNav,
        Breadcrumb,
        Pagination,
        LogItem
    },
    computed: {
        routesNav: function(){
          return [
                {name:'Usuarios'},
                {name:'Editar usuario' , path: '/usuarios/' + this.$route.params.idUser + '/editar-usuario'},
                {name:'Asignar mapas' , path: '/usuarios/' + this.$route.params.idUser + '/asignar-mapas'},
                {name:'Historial de cambios' , path: '/usuarios/' + this.$route.params.idUser + '/historial-de-cambios'},
                {name:'Eliminar usuario' , path: '/usuarios/' + this.$route.params.idUser + '/eliminar-usuario'}
              ]
        },
        routesBread: function(){
          return [
              {name:'Usuarios', path: '/usuarios' },
              {name: this.user.firstname + ' ' + this.user.lastname}
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
            this.getUserLogs(this.currentPage);
        },
        getGroupByID: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/' + this.user.id_group)
            .then(function (response) {
                that.userGroup = response.data;
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getUserLogs: function(page) {
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/users/' + this.$route.params.idUser + '/logs?page=' + page)
            .then(function (response) {
                that.logs = response.data;
                response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getAssignedMaps: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/users/' + this.$route.params.idUser + '/maps') //Sin paginación
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
        //Verificar permisos
        this.$store.getters.getPermissions.users? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getUser();
        this.getAssignedMaps()
    },
    watch: {
        'user': function (newValue, oldValue) {
            if(newValue.id_group){
                this.getGroupByID();
            }
        },
    }
}
</script>

<style lang="scss" scoped>

.permissions{
    display: grid;  
    grid-template-columns: auto auto auto auto; 
    grid-column-gap: 2em; 
    justify-content: start;
}

@media only screen and (max-width: 640px) {
    .permissions{
        grid-template-columns: auto auto; 
    }
}

</style>
