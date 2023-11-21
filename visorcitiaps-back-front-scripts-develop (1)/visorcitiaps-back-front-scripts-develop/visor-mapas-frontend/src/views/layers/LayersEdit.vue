<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Editar capa</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="name">Nombre</label>
        <input id="name" v-model="layer.name" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre de la capa']]"/>

        <label for="category">Categoría</label>
        <select id="category"
        v-model="layer.id_category"
        name="categoría de la capa" :class="{'error' : errors.has('categoría de la capa') }" v-validate="categoriesIncluded">
            <option
            value="-1"
            disabled
            >
                Selecciona una categoría
            </option>
            <option
            v-for="item in categories"
            v-bind:key="item.index"
            v-bind:value="item.id">
            {{item.name}}
            </option>
        </select>

        <label for="desc">Descripción</label>
        <textarea id="desc" v-model="layer.desc" placeholder="Ingresa la descripción de la capa"></textarea>
        <div v-if="layer.provider.name=='file'">
        <label for="filename">Nombre del archivo</label>
        <input id="filename" v-model="layer.provider.geoserverdata.filename" type="text"
        name="nombre del archivo" :class="{'error' : errors.has('nombre del archivo') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre del archivo') ? errors.first('nombre del archivo') : 'Nombre de la capa en GeoServer']]"/>

        <!--
        <label for="coordinatesSystem">Sistema de coordenadas</label>
        <input id="coordinatesSystem" v-model="layer.provider.geoserverdata.coordinates_system" type="text"
        name="sistema de coordenadas" :class="{'error' : errors.has('sistema de coordenadas') }" v-validate="'required'"
        :placeholder="[[errors.has('sistema de coordenadas') ? errors.first('sistema de coordenadas') : 'Sistema de coordenadas de la capa']]"/>
        -->
        
        <label for="workspace">Espacio de trabajo</label>
        <input id="workspace" v-model="layer.provider.geoserverdata.workspace" type="text"
        name="espacio de trabajo" :class="{'error' : errors.has('espacio de trabajo') }" v-validate="'required'"
        :placeholder="[[errors.has('espacio de trabajo') ? errors.first('espacio de trabajo') : 'Espacio de trabajo de la capa en GeoServer']]"/>
        </div> 
        <div v-else>
        <label for="name">Url</label>
        <input id="url" v-model="layer.provider.url" type="text"
        name="url" :class="{'error' : errors.has('url') }" v-validate="'required'"
        :placeholder="[[errors.has('url') ? errors.first('url') : 'Ingresa la url de la capa']]"/>
        </div>

        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()"
            v-shortkey="['shift', 'z']"
            @shortkey="onClickGoBack"
            >Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()"
            v-shortkey="['shift', 'enter']"
            @shortkey="onClickGoNext"
            >Actualizar capa</button>
        </div>

    </div>

    <div class="container-list"> 
        
    </div>
 
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {restLayer} from '@/components/mixin/restLayer.js'
import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [restLayer, restCategory],
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
              {name:'Capas', path: '/capas' },
              {name: this.layer.name, path: '/capas/' + this.$route.params.idLayer},
              {name: 'Editar capa'}
              ]
        },
        categoriesIncluded: function(){
            let list = 'included:';
            for (let i = 0; i < this.categories.length; i++) {
                list = list + this.categories[i].id.toString() + ',';
            }
            return list;
        }
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/' + this.$route.params.idLayer)
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    console.log('OK')
                    that.putLayer();
                } else {
                    console.log('NOPE')
                }
            })
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getCategoriesNoPages();
        this.getLayer();
    }
}
</script>

<style lang="scss" scoped>

</style>
