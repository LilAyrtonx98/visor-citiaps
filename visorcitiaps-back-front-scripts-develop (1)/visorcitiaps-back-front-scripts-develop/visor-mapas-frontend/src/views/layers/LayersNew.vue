<template>
<main class="container-1-1">

    <div class="container-header">
        <Breadcrumb
        v-bind:links="routesBread"
        >
        </Breadcrumb>
        <h2>Nueva capa</h2>
    </div>

    <div class="container-action">
        
    </div>

    <div class="container-body">
        <label for="name">Nombre</label>
        <input id="name" v-model="layer.name" type="text"
        name="nombre" :class="{'error' : errors.has('nombre') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre') ? errors.first('nombre') : 'Ingresa el nombre de la capa']]"/>

        <label for="category">Categoría</label>
        <br>
        <select id="category"
        v-if="categories"
        v-model="layer.id_category"
        name="categoría de la capa" :class="{'error' : errors.has('categoría de la capa') }" v-validate="categoriesIncluded">
            <option
            value="-1"
            disabled
            >
                {{errors.has('categoría de la capa') ? 'Selecciona un categoría para la capa' : 'Selecciona una categoría'}}
            </option>
            <option
            v-for="item in categories"
            v-bind:key="item.index"
            v-bind:value="item.id">
            {{item.name}}
            </option>
        </select>
        <span v-else>No hay categorías que seleccionar para la capa</span>
        <br>

        <label for="desc">Descripción</label>
        <textarea id="desc" v-model="layer.desc" placeholder="Ingresa la descripción de la capa"></textarea>

        <label for="creationmode">Nueva capa a partir de</label>
        <div id="creationmode">
            <input type="radio" id="shapefile" value="file" v-model="layerCreationMode">
            <label for="shapefile" style="font-size: 0.9em; font-weight: normal;"> Solo subir SHP a partir de .Zip</label>
            <br>
            <input type="radio" id="geoserver" value="geoserver" v-model="layerCreationMode">
            <label for="geoserver" style="font-size: 0.9em; font-weight: normal;"> Servicio OWS</label>
            <br>
            <input type="radio" id="arcgis" value="arcgis" v-model="layerCreationMode">
            <label for="arcgis" style="font-size: 0.9em; font-weight: normal;"> Servicio ESRI</label>
            <br>
        </div>

        <label v-if="layerCreationMode=='file'" style="font-size: 0.9em; font-weight: normal;">La capa se registrará en el visor y se subirá a GeoServer</label>
        <label v-if="layerCreationMode=='geoserver'" style="font-size: 0.9em; font-weight: normal;">La capa se registrará en el visor de acuerdo a los datos existentes en GeoServer</label>
        <label v-if="layerCreationMode=='arcgis'" style="font-size: 0.9em; font-weight: normal;">La capa se registrará en el visor de acuerdo a los datos existentes en Arcgis</label>
        <br>

        <template v-if="layerCreationMode=='file'">
            <label for="file-button" style="display: block;">Subir archivo de capa</label>
            <div id="file-button" class="search-tools">
                <label for="file" class="btn btn-secondary file-upload">
                    Seleccionar
                    <input id="file" accept=".zip" type="file" ref="file" v-on:change="onHandleFileChange()"/>
                </label>
                <span>{{uploadFileMessage}}</span>
            </div>
        </template>

        <label v-if="layerCreationMode=='file'" for="filename">Nombre del archivo</label>
        <input v-if="layerCreationMode=='file'" id="filename" v-model="layer.provider.geoserverdata.filename" type="text"
        name="nombre del archivo" :class="{'error' : errors.has('nombre del archivo') }" v-validate="'required'"
        :placeholder="[[errors.has('nombre del archivo') ? errors.first('nombre del archivo') : 'Nombre del archivo']]"/>

        <label v-if="layerCreationMode!='file'">Url</label>
        <input v-if="layerCreationMode!='file'" v-model="layer.provider.url" placeholder="Ingresa la url del servicio">
        
        
        <label v-if="layerCreationMode=='shapefile'" for="coordinatesSystem">Sistema de coordenadas</label>
        <input v-if="layerCreationMode=='shapefile'" id="coordinatesSystem" v-model="layer.provider.geoserverdata.coordinatessystem" type="text"
        name="sistema de coordenadas" :class="{'error' : errors.has('sistema de coordenadas') }" v-validate="'required'"
        :placeholder="[[errors.has('sistema de coordenadas') ? errors.first('sistema de coordenadas') : 'Sistema de coordenadas de la capa']]"/>    
        <p style="font-size: 0.9em;" v-if="layerCreationMode=='shapefile'">Es necesario especificar el sistema de coordenadas de la capa si se quiere agregar a GeoServer. Si no es correcto, la capa no será procesada correctamente.</p>
        
        <label v-if="layerCreationMode=='geoserver'" for="workspace">Espacio de trabajo</label>
        <input v-if="layerCreationMode=='geoserver'" v-model="layer.provider.geoserverdata.workspace" placeholder="Espacio de trabajo de la capa en GeoServer">

        <label v-if="layerCreationMode=='geoserver'" for="datastore">Almacen de datos</label>
        <input v-if="layerCreationMode=='geoserver'" v-model="layer.provider.geoserverdata.datastore" placeholder="Almacen de datos de la capa en GeoServer">


        <div class="container-action-left" style="justify-content: start;">
            <button class="btn btn-secondary" v-on:click="onClickGoBack()">Volver</button>
            <button class="btn btn-secondary" v-on:click="onClickGoNext()">Crear capa</button>
        </div>

    </div>

    <div class="container-list"> 
        
    </div>
 
</main>
</template>

<script>
// @ is an alias to /src
import Breadcrumb from '@/components/Breadcrumb.vue'

import {formatter} from '@/components/mixin/formatter.js'
import {restLayer} from '@/components/mixin/restLayer.js'
import {restCategory} from '@/components/mixin/restCategory.js'

export default{
    name: 'test',
    mixins: [formatter, restLayer, restCategory],
    data(){
        return{
            file: '',
            layerCreationMode: 'file',
        }
    },
    components:{
        Breadcrumb
    },
    computed: {
        routesBread: function(){
          return [
              {name:'Capas', path: '/capas' },
              {name: 'Nueva capa'}
              ]
        },
        categoriesIncluded: function(){
            let list = 'included:';
            for (let i = 0; i < this.categories.length; i++) {
                list = list + this.categories[i].id.toString() + ',';
            }
            return list;
        },
        uploadFileMessage: function() {
            if (this.file == ''){
                return 'No ha seleccionado un archivo .zip';
            } else {
                return this.file.name;
            }   
        },
    },
    props: [],
    methods:{
        onClickGoBack: function() {
            window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/capas/')
        },
        layerCategoryNameByID: function(idLayer){
            var category = this.categories.find(function(element) {
                    return element.id == idLayer;
                });
            return category.name;
        },
        onClickGoNext: function() {
            var that = this;
            this.$validator.validateAll()
            .then(function (response) {
                if(response) {
                    if(that.layerCreationMode == 'file') {
                        console.log('file upload selection');
                        that.postFile();
                    } else {
                        console.log('otro');
                        that.layer.provider.name=that.layerCreationMode;
                        that.layer.provider.geoserverdata.filename = that.cleanString(that.layer.provider.geoserverdata.filename.toLowerCase());
                        that.layer.provider.geoserverdata.workspace = that.cleanString(that.layer.provider.geoserverdata.workspace.toLowerCase());
                        that.layer.provider.geoserverdata.datastore = that.cleanString(that.layer.provider.geoserverdata.datastore.toLowerCase());
                        console.log('1 put layer', that)
                        that.postLayer();
                    }
                } else {
                }
            })
        },
        onHandleFileChange: function() {
            this.file = this.$refs.file.files[0];
            this.layer.provider.geoserverdata.filename = this.file.name.split('.')[0];
        },
        postFile: function() {
            console.log('layer 1', this.layer)
            this.layer.provider.name= this.layerCreationMode;
            this.layer.provider.geoserverdata.filename = this.cleanString(this.layer.provider.geoserverdata.filename.toLowerCase());
            this.layer.provider.geoserverdata.workspace = process.env.VUE_APP_DEFAULT_WORKSPACE;
            this.layer.provider.geoserverdata.datastore = process.env.VUE_APP_DEFAULT_DATASTORE;
            
            var that = this;
            let formData = new FormData();
            formData.append('file', this.file);
            formData.append('coordinateSystem', this.layer.provider.geoserverdata.coordinatessystem);
            formData.append('workspace', process.env.VUE_APP_DEFAULT_WORKSPACE);
            formData.append('datastore', process.env.VUE_APP_DEFAULT_DATASTORE);

            if(!this.layer.provider.geoserverdata.filename.includes('.zip')) {
               this.layer.provider.geoserverdata.filename = this.layer.provider.geoserverdata.filename + '.zip';
            }
            formData.append('filename', this.layer.provider.geoserverdata.filename);

            this.$http.post(this.$store.getters.backendurl + 'api/v1/uploads/single',
            formData,
            {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            })
            // Este postLayer debe generarse aparte (Con una existente en Geoserver) debido a que la url de la capa
            // es generada luego de ser subida al servidor (Geoserver local)
            // .then(function(){
            //     console.log('2 put layer', that)
            //     that.postLayer();
            //     that.layer.provider.geoserverdata.filename = that.layer.provider.geoserverdata.filename.replace('.zip', '')
            // })
            .then(function(){
                console.log('put layer', that)
                console.log('Layer in geoserver');
                that.$router.push('/capas');
            })
            .catch(function(){
                that.$router.push('/capas');
            });
            
        }
    },
    updated () {

    },
    created () {
        this.$store.getters.getPermissions.layers? true:this.$router.push('/pagina-no-encontrada')
    },
    mounted () {
        this.getCategoriesNoPages()
    },
    watch: {
        // whenever question changes, this function will run
        layerCreationMode: function (newValue, oldValue) {
            if(newValue == 'geoserver') {
                this.filename = '';
            } else {
                if(this.file != '') {
                    this.layer.provider.geoserverdata.filename = this.file.name.split('.')[0];
                }
            }
        }
  },
}
</script>

<style lang="scss" scoped>

.file-upload{
    display: block;
    cursor: pointer;
    font-weight: bold;
    margin-top: 0em;
    font-size: 0.8em;
    opacity: 0.8;
    white-space: nowrap;
    text-align: center;
    input{
        display: none;
        /*
        background: white;
        border: none;
        font-size: 1em;
        border-radius: 0px;
        width: 100%;
        padding-left: 0em;
        */
    }
}
.search-tools{
    display: grid;
    grid-template-columns: min-content auto;
    grid-column-gap: 0.5em;
    align-items: end;
}

</style>
