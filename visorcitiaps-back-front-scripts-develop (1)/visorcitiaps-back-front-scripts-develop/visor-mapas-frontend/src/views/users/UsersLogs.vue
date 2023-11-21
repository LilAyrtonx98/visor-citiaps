<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Historial de cambios</h2>
    </div>
    
    <div class="container-body">
        <LogItem
        v-bind:logs="logs"
        v-bind:showInfo="true"
        v-bind:fullMode="true"
        @itemSelected="responseSelectedItem">
        </LogItem>

        <Pagination
        v-show="totalPages!=0"
        v-bind:totalPages="totalPages"
        v-bind:currentPage="currentPage"
        v-bind:pageSize="pageSize"
        v-on:pagechanged="pageChangedClick"
        >
        </Pagination>
    </div>

    <div class="container-list">
        <strong>Recurso del registro seleccionado</strong>
        <br>
        <div v-if="selectedItem" >
            <button class="btn btn-secondary" v-clipboard:copy="JSON.stringify(selectedItem)">
                <font-awesome-icon icon="copy"></font-awesome-icon>
                Copiar JSON
            </button>  
        </div>
        <span v-else>No has seleccionado un recurso de un registro anterior</span>
        <json-view 
        v-if="selectedItem"
        :data="selectedItem"
        rootKey="Usuario"
        :maxDepth="3"
        />
    </div>

</main>
</template>

<script>
// @ is an alias to /src
import SideNav from '@/components/SideNav.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'
import Pagination from '@/components/Pagination.vue'
import LogItem from '@/components/LogItem.vue'
import { JSONView } from "vue-json-component";

import {formatter} from '@/components/mixin/formatter.js'
import {pagination} from '@/components/mixin/pagination.js'
import {restUser} from '@/components/mixin/restUser.js'

export default{
    name: 'test',
    mixins: [formatter, pagination, restUser],
    data(){
        return{
            message: '',
            user: '',
            logs: [],
            selectedItem: null
        }
    },
    components:{
        SideNav,
        Breadcrumb,
        Pagination,
        LogItem,
        "json-view": JSONView
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Usuarios', path: '/usuarios' },
              {name: this.user.firstname + ' ' + this.user.lastname, path: '/usuarios/' + this.$route.params.idUser},
              {name: 'Historial de cambios'}
              ]
        },
    },
    props: [],
    methods:{
        pageChangedClick: function (page){
            this.currentPage = page;
            this.getLogs(this.currentPage);
        },
        getLogs: function(page) {
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/logs/users/' + this.$route.params.idUser + '?page=' + page)
            .then(function (response) {
                that.logs = response.data
                response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        responseSelectedItem: function(item) {
            if(item.new_resource == "")
            {
                this.selectedItem = null;
                return;
            }
            this.selectedItem = JSON.parse(item.new_resource);
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
        this.getUser();
        this.getLogs(this.currentPage);
    }
}
</script>

<style lang="scss">

#json-view{
    .properties{
        display: none;
    }
}

</style>
