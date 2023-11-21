<template>
<div class="list-item-box">
    <div v-for="(item) in list">
        <div class="list-item-item">
            <span v-if="title2" v-bind:key="item.index" class="list-item-name" >{{item[title]+" "+item[title2]}}</span>
            <span v-else v-bind:key="item.index" class="list-item-name" >{{item[title]}}</span>
            <span v-if="subTitle2" v-bind:key="item.index" class="list-item-sub-1" >{{item[subTitle][subTitle2]}}</span>
            <span v-else v-bind:key="item.index" class="list-item-sub-1" >{{item[subTitle]}}</span>
            <span v-bind:key="item.index" class="list-item-sub-3">
                <button v-bind:key="item.index" v-if="!item[isAssingned]" class="btn btn-secondary" v-on:click="assing(item)">{{assingText}}</button>
            </span>
            <span v-bind:key="item.index" class="list-item-separator"></span>
        </div>
    </div>
</div>
</template>

<script>
export default {
    name: 'assingBox',
    props: ['list', 'title', 'title2', 'subTitle', 'subTitle2','isAssingned', 'assingText'],
    methods:{
        assing(item){
            this.$emit('assingMethod', item)
        }
    }
}
</script>