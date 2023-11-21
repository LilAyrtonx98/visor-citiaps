<template>
    <div class="list-item-alt">
        <div class="list-item-info-alt">
            <strong>{{layer.name}}</strong>
            <div class="list-item-info-alt">
                <div class="option" style="padding-left: 0em;">
                    <button class="btn-collapsible" @click="showList=!showList">
                        <font-awesome-icon v-if="!showList" icon="chevron-down" style="font-size: 2em;"></font-awesome-icon>
                        <font-awesome-icon v-else icon="chevron-up" style="font-size: 2em;"></font-awesome-icon>
                    </button>
                </div>
            </div>
        </div>

        <div class="items-list-collap" v-show="showList">
            <img :src="legendSource" width="100%"> 
        </div>
        
    </div>
</template>

<script>
import {formatter} from '@/components/mixin/formatter.js'

export default {
    props:['layer'],
    mixins:[formatter],
    data(){
        return{
            showList: false,
        }
    },
    methods:{
        getLegendSource: function() {
            return legendSource
        }
    },
    updated(){
    },
    computed:{
        legendSource: function(){
            if(this.layer.provider.name == 'geoserver'){
                console.log('GEOSERVER vue getLegendSource');
                var url = this.layer.provider.parsed_url.host +':'+this.layer.provider.parsed_url.port + '/geoserver/wms?'
                var params =
                '&REQUEST=GetLegendGraphic' +
                '&VERSION=1.1.1' + 
                '&FORMAT=image/png' + 
                '&WIDTH=100' +
                '&HEIGHT=20' + 
                '&LAYER=' +  this.layer.provider.geoserverdata.workspace + ':' + this.layer.provider.geoserverdata.datastore +
                '&LAYER_WORK=' + this.layer.provider.geoserverdata.workspace
                return url+params
            }else{
                console.log('ARGIS vue getLegendSource');
            }
        }
    }
}
</script>

<style scoped lang="scss">
p{
    margin: 0.5em 0;
}
.list-item-alt{
    //border-bottom: 1px solid rgba(0,0,0,0.5);
    color: #4D4D4D;
    font-size: 0.9em;
    padding: 0.25em 1em;
    margin-top: 0.5em;
    margin: 0.25em 0em;
}
.bottom-menu-fixed{
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
}
.options-fixed{
    display: inherit;
    align-items: start;
    position: relative;
    cursor: pointer;
}

.item{
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    background-color: #FCFCFC;
    padding: 10px;
    border-bottom: 1px solid #DBDBDB;
    color: #4D4D4D;
    font-size: 1em;
}

.training-title{
    margin: 0.5em;
}

.training-text{
    display: block;
    margin: 0.5em;
}

.list-item-info-alt{
  display: flex;
  justify-content: space-between;
  align-items: center;
  //margin-top: 0.5em; 
  //margin-bottom: 0.5em;
}

.items-list-collap{
    //padding-left: 1em;
    div{
        padding: 0.125em 0em;
    }
}

.btn-collapsible{
    border: none;
    background: transparent;
    color: #036016;
}

</style>

