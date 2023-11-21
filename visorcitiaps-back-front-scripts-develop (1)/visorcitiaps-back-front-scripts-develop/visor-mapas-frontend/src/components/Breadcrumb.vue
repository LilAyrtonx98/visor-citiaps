<template>

  <div class="breadcrumb-alt">
    <template v-for="link in links">

      <router-link v-bind:key="link.index" v-if="link.path" :to="link.path" :title="link.name">
        <span>{{link.name}}</span>
      </router-link>
      <span class="router" v-bind:key="link.index" v-else :title="link.name">{{link.name}}</span>
      <span class="separator" v-bind:key="link.index" v-if="link.path">></span>

    </template>
  </div>

</template>

<style lang="scss" scoped>
.breadcrumb-alt{
  display: grid;
  grid-template-columns: auto min-content auto min-content auto min-content auto min-content;
  justify-content: start;
  .separator{
    font-weight: bold;
    padding-right: 0.25em;
    padding-left: 0.25em;
  }
  a, .router{
    font-weight: bold;
    white-space: nowrap;

    overflow: hidden;
    text-overflow: ellipsis;
  }
}

.breadcrumb-alt .active:hover{
  background-color: transparent;
}
/*
.breadcrumb-alt .router:hover{
  background-color: blue;
  color: white;
}
*/

// 1000px
///////////////////////
/*
@media all and (max-width: 420px) {
	.breadcrumb-alt{
    grid-template-columns: auto min-content;
    justify-content: start;
    a, .router{
      font-weight: bold;
      white-space: nowrap;
      padding-right: 0em;
      padding-left: 0em;

      overflow: hidden;
      text-overflow: ellipsis;
    }
      a:first-child{
      padding-left: 0em;
    }
    a:last-child{
      padding-right: 0em;
    }
  }
}
*/
</style>


<script>
export default{
  props:['links'],
  data(){
    return {
        
    }
  },
  computed: {

  },
  methods: {
      getUrl: function(link) {
          return link.path;
        },
  },
  created(){
  
  },
  watch: {

  }
}
</script>