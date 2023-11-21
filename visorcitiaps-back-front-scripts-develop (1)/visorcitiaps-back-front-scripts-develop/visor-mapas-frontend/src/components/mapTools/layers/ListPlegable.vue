<template>
    <div class="list-item-alt">
        <div class="list-item-info-alt">
            <strong>{{category.name}}</strong>
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
            <div v-for="(layer, j) in category.layers" v-bind:key="j" style="overflow:auto;">
                <input type="checkbox" :id="j" v-model="layer.selected" v-on:change="clickLayer(j, layer.selected)">
                <label  style="font-size: 0.9em; font-weight: normal;" :disabled="true" >{{layer.name}}</label>
                <font-awesome-layers @click="setLayerMenu(j)" class="fa-lg" style="float:right; padding:0px 20px 0px 20px; cursor:pointer">
                    <font-awesome-icon icon="ellipsis-h"/>
                </font-awesome-layers>
                <ul v-if="menuId==j" class="layerMenu" v-click-outside="hideMenu" @click="clickShowAttributes(j)">
                        Ver tabla de atributos
                </ul>
                <br>
            </div>   
        </div>
        
    </div>
</template>

<script>
import {formatter} from '@/components/mixin/formatter.js'
import vClickOutside from 'v-click-outside'
export default {
    props:['index','category'],
    mixins:[formatter],
    data(){
        return{
            showList: false,
            selected: false,
            menuId: -1
        }
    },
    directives: {
      clickOutside: vClickOutside.directive
    },
    methods:{
        openModal: function(){
            this.$emit("showModal", this.index)
        },
        clickLayer: function(layerIndex, isSelected){
            if(isSelected) {
                this.$emit("selectLayer", this.index, layerIndex)
            } else {
                this.$emit("removeLayer", this.index, layerIndex)
            }
        },
        setLayerMenu: function(id){
            console.log(id)
            if(this.menuId==-1) this.menuId=id
            else this.menuId=-1
        },
        hideMenu: function(){
            this.menuId=-1
        },
        clickShowAttributes: function(layerIndex){
            this.menuId=-1
            this.$emit("layerAttributes", this.index, layerIndex)
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

