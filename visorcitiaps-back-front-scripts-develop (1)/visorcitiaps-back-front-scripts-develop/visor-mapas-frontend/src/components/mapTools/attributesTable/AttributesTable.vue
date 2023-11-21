<template>
    <div v-show="isActive" class="attributesContainer">
        <div v-if="!isFixed" class="attributesHeader">
            <div v-if="layer" style="font-weight: bold;">
                <template v-if="layer.provider.name=='arcgis'">
                    <p>{{layer.name}} </p>
                    <select v-if="subLayers.length>1" v-model="layerId">
                        <option v-for="l in subLayers" :value="l.id">{{l.name}}</option>
                    </select>
                </template>
                <template v-else ><p>{{layer.name}}</p></template>
            </div>

            <div>

            <button class="btn btn-info" v-if="recordsPerPage<totalRecords" :disabled="firstPage" v-on:click="prevPage()">
                <font-awesome-icon icon="chevron-left"></font-awesome-icon>
            </button>
            <button class="btn btn-info" v-if="recordsPerPage<totalRecords" :disabled="lastPage"  v-on:click="nextPage()">
                <font-awesome-icon icon="chevron-right"></font-awesome-icon>
            </button><!-- 
            <button class="btn btn-info" v-on:click="isHidden = !isHidden">
                <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
                <font-awesome-icon v-else icon="plus"></font-awesome-icon>
            </button> -->
            <button class="btn btn-danger" v-on:click="exitAttributes()">
                <font-awesome-icon icon="times"></font-awesome-icon>
            </button>
            </div>
        </div>
        <div class="attributesBody" v-show="!isHidden">
            <div style="overflow-x: scroll; height:300px; overflow-y: scroll">
            <table class="fixed_headers">
                <thead>
                    <tr>
                        <th v-for="h in headers">{{h}}</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="d in data">
                        <td v-for="h in headers">{{d[h]}}</td>
                    </tr>
                </tbody>
            </table></div>
        </div>
    </div>
</template>
<script>
import {vuex} from '@/components/mixin/vuex.js'
export default {
    data(){
        return{
            data:[],
            headers:[],
            isHidden: false,
            recordsPerPage:1000,
            offset:-1,
            totalRecords:0,
            layer: null,
            subLayers:[],
            layerId: null
        }
    },
    props: ['isFixed'],
    mixins:[vuex],
    methods:{
        exitAttributes: function(){
            this.selectMapTool('attributesTable', false)
        },
        prevPage: function(){
            console.log("pag anterior")
            this.offset=this.offset-this.recordsPerPage
        },
        nextPage: function(){
            console.log("pag siguiente")
            this.offset=this.offset+this.recordsPerPage
        },
        getFeatures: function(){
            var that=this
            this.data=[]
            if(this.layer.provider.name=='arcgis'){
                var req = new XMLHttpRequest();
                req.open('GET', this.layer.provider.url+'/'+this.layerId+'/query?where=OBJECTID>'+this.offset+'&outFields=*&f=pjson', false);
                req.withCredentials = false;
                req.send();

                if (req.status == 200){
                    let temp=JSON.parse(req.response);
                    for(let i=0; i<temp.features.length; i++) {
                        that.data.push(temp.features[i].attributes)
                    }
                    that.headers=Object.getOwnPropertyNames(temp.fieldAliases)
                }else{
                    console.log("Fail Feature argis query");
                }

                // DEPRECATED CORS ERROR needed withCredentials = false
                // this.$http.get(this.layer.provider.url+'/'+this.layerId+'/query?where=OBJECTID>'+this.offset+'&outFields=*&f=pjson')
                // .then(function(response){
                //     let temp=response.data
                //     for(let i=0; i<temp.features.length; i++) {
                //         that.data.push(temp.features[i].attributes)
                //     }
                //     that.headers=Object.getOwnPropertyNames(temp.fieldAliases)
                // })
            }
            else{
                var url=this.layer.provider.url+'&count='+this.recordsPerPage+'&startIndex='+this.offset+'&outputFormat=application%2Fjson'


                var req = new XMLHttpRequest();
                req.open('GET', url, false);
                req.withCredentials = false;
                req.send();

                if (req.status == 200){
                    console.log("Success Feature geo query");
                    let temp=JSON.parse(req.response)
                    that.headers=Object.getOwnPropertyNames(temp.features[0].properties)
                    for(let i=0; i<temp.features.length; i++) {
                        that.data.push(temp.features[i].properties)
                        
                    }
                }else{
                    console.log("Fail Feature geo query");
                }

                // this.$http.get(url)
                // .then(function(response){
                //     let temp=response.data
                //     that.headers=Object.getOwnPropertyNames(temp.features[0].properties)
                //     for(let i=0; i<temp.features.length; i++) {
                //         that.data.push(temp.features[i].properties)
                        
                //     }
                // })
            }
        }
    },
    computed: {
        isActive: function() {
            return this.isFixed? this.isMobile : this.mapToolValues('attributesTable').active && !this.isMobile
        },
        firstPage: function(){
            if (this.offset==0) return true
            else return false
        },
        lastPage: function(){
            if(this.offset+this.recordsPerPage>this.totalRecords) return true
            else return false
        }
    },
    watch:{
        'mapToolExecuted': function(){
            if(this.mapToolValues('attributesTable').action == 'sendLayerInfo'){
                this.data=this.mapToolValues('attributesTable').data.layerData
                this.headers=this.mapToolValues('attributesTable').data.layerProperties
                this.layer=this.mapToolValues('attributesTable').data.layer
            }

            if(this.mapToolValues('attributesTable').action == 'paginationInfo'){
                this.totalRecords=this.mapToolValues('attributesTable').data.totalRecords
                this.layer=this.mapToolValues('attributesTable').data.layer
                this.offset=this.mapToolValues('attributesTable').data.offset
            }
        },
        'layer': function(){
            if(this.layer.provider.name=='arcgis'){
                this.layerId=this.mapToolValues('attributesTable').data.layerId
                this.subLayers=this.mapToolValues('attributesTable').data.subLayers
            }
            this.getFeatures()
        },
        'offset': function(){
            this.getFeatures()
        },
        'layerId': function(){
            this.offset=0
            this.getFeatures()
        }
    }
}
</script>