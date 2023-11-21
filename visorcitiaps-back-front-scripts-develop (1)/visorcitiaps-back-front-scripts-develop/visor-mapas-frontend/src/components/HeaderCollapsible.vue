<template>
  <header class="header">

    <span style="color: white; align-self: center; font-weight: bold; font-size: 1.2em; white-space: nowrap;">Visor de mapas</span>

    <nav id="nav" class="navigation">
      <router-link v-for="link in links" v-bind:key="link.index" :to="link.path">{{link.name}}</router-link>
    </nav>
    
    <button id="collapsible-button" v-on:click="onClickSelectCollapsible()" class="collapsible-button">
      <font-awesome-icon v-if="!isActive" icon="chevron-down" style="font-size: 2em;"></font-awesome-icon>
      <font-awesome-icon v-else icon="chevron-up" style="font-size: 2em;"></font-awesome-icon>
    </button>
    
    <div id="collapsible-content" class="collapsible-content">
      <nav id="nav" class="navigation-toggle">
        <router-link v-for="link in links" v-bind:key="link.index" v-on:click.native="onClickSelectCollapsible()" :to="link.path">{{link.name}}</router-link>
        <router-link
        v-if="account"
        v-on:click.native="onClickSelectCollapsible()"
        :to="'/cuenta/mi-cuenta'"
        >
          Mi cuenta
        </router-link>
        <router-link 
        v-if="account"
        v-on:click.native="onClickSelectCollapsible()"
        :to="'/cuenta/cerrar-sesion'"
        >
          Cerrar sesión
        </router-link>
      </nav>
    </div>

    <UserDropdown class="dropdown-user" v-if="account"></UserDropdown>

  </header>
</template>

<style lang="scss" scoped>
/*
  style.scss debe tener definido body>header de esta forma: (grid de tres columnas)

  body>header {
    display: grid;
    grid-template-columns: 1fr 1fr min-content;
    justify-content: space-between;
  }
*/
.navigation {
  display: flex;
  align-items: flex-end;
  justify-self: flex-end;
}
.navigation-toggle{
    display: none;
  }
.collapsible-content {
  display: none;
}
.collapsible-button {
  display: none;
}

@media only screen and (max-width: 640px) {
  header{
    align-content: center;
  }
  .dropdown-user{
    display: none;
  }
  .navigation {
    display: none;
  }
  .navigation-toggle{
    display: grid;
    grid-template-columns: 1fr;
  }
  .collapsible-content {
    grid-column-start: 1;
    grid-column-end: 3;
    display: grid; //display: grid; causa problemas de espacio. Se cambia mediante javascript al hacer click en el botón
    grid-template-columns: 1fr;
    max-height: 0;
    width: 100%;
    overflow: hidden;
    transition: max-height 0.2s ease-out;
    background-color: transparent;
    text-align: center;
  }
  .collapsible-button {
    display: inline;
    background-color: transparent;
    color: white;
    cursor: pointer;
    border: none;
    outline: none;
    font-size: 1em;
    text-align: end;
    justify-self: flex-end;
    align-self: center;
  }
 
}
</style>

<script>
import UserDropdown from '@/components/UserDropdown.vue'

export default {
  name: 'NavBarCollapsible',
  data(){
    return{
      isActive: false
    }
  },
  components:{
      UserDropdown
  },
  computed: {

  },
  props:['links', 'account'],
  methods:{
    onClickSelectCollapsible: function() {
      let content = document.getElementById("collapsible-content");   
      
      this.isActive = !this.isActive;

      if (content.style.maxHeight){
        content.style.maxHeight = null;
        //content.style.display = "none";
      } else {
        //content.style.display = "grid";
        content.style.maxHeight = content.scrollHeight + "px";
      } 
    }
  },
  updated () {

  },
  created () {

  },
  mounted () {

  }
}
</script>
