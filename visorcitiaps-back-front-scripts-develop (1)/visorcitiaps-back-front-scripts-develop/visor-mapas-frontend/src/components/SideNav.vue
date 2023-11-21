<template>
  <div class="nav-side">

    <template v-for="link in links">

      <button
      class="btn btn-danger nav-button"
      v-bind:class="{active : isActive(link.path)}"
      v-bind:key="link.index"
      v-if="link.name.includes('Eliminar')"
      v-on:click="onClickGoTo(link.path)"

      v-shortkey="['shift', 'ctrl', link.name.toLowerCase().charAt(0)]"
      @shortkey="onClickGoTo(link.path)"
      >
        {{link.name}}
      </button>

      <button
      class="btn btn-secondary nav-button"
      v-bind:class="{active : isActive(link.path)}"
      v-bind:key="link.index"
      v-else-if="link.path"
      v-on:click="onClickGoTo(link.path)"

      v-shortkey="['ctrl', link.name.toLowerCase().charAt(0)]"
      @shortkey="onClickGoTo(link.path)"
      >
        {{link.name}}
      </button>

      <span v-bind:key="link.index" v-else>{{link.name}}</span>

    </template>
    
  </div>
</template>

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
      onClickGoTo: function(link) {
        //Ya está aquí
        if(this.$route.fullPath.includes(link))
        {
          return;
        }
        this.$router.push(link);
      },
      isActive: function(link) {
        if(this.$route.fullPath.includes(link))
        {
          return true;
        }
        return false;
      },
      
  },
  created(){
  
  },
  watch: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.nav-side{
  display: grid;
  grid-template-columns: 1fr;
  grid-row-gap: 0.5em;
  white-space: nowrap;
  justify-content: normal;
  grid-auto-rows: max-content;
  justify-content: space-evenly;
  justify-self: end;
	width: 90%;
}

.nav-button{
  
}

.active{
  background: #f7e200;
}

@media only screen and (max-width: 640px) {

  .nav-side{
    grid-template-columns: 1fr;
    grid-row-gap: 0.25em;
    justify-content: space-evenly;
    justify-self: start;
	  width: 100%;
  }

  .nav-button{
    
  }

}
</style>
