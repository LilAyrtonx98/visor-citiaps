<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Nuevo geoproceso</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="name">Nombre</label>
        <input id="name" v-model="geo.name" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre del geoproceso']]"/>

        <label for="desc">Descripción</label>
        <textarea id="desc" v-model="geo.desc" placeholder="Ingresa la descripción del geoproceso"></textarea>

        <label for="url">URL</label>
        <input id="url" v-model="geo.geo_url" type="text"
        name="url" :class="{'error' : errors.has('url') }" v-validate="'required'"
        :placeholder="[[errors.has('url') ? errors.first('url') : 'Ingresa la url del geoproceso']]"/>

        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Crear geoproceso</button>
        </div>

    </div>

    <div class="container-list"> 
        
    </div>
 
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'
import {restGeoprocessing} from '@/components/mixin/restGeoprocessing.js'

export default{
    name: 'test',
    mixins: [restGeoprocessing],
    data(){
        return{

        }
    },
    components:{
        Breadcrumb
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Geoprocesos', path: '/geoprocesos' },
              {name: 'Nuevo geoproceso'}
              ]
        }
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/')
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    console.log('OK')
                    that.postGeoprocessing();
                } else {
                    console.log('NOPE')
                }
            })
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.geo? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {

    }
}
</script>

<style lang="scss" scoped>

</style>
