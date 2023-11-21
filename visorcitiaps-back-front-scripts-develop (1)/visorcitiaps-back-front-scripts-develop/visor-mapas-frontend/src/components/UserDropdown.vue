<template>
  <div class="dropdown">

    <a class="dropbtn">Cuenta</a>

    <div class="dropdown-content">
      <router-link 
      :to="'/cuenta/mi-cuenta'"
      >
        Mi cuenta
      </router-link>
      <router-link 
      :to="'/cuenta/cerrar-sesion'"
      >
        Cerrar sesión
      </router-link>
    </div>

  </div>
</template>

<script>
export default {
  name: 'UserDropdown',
  computed: {
  },
  methods: {
  
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.dropdown {
  align-self: flex-end;
  padding: 0.8em 0.5em;
  font-size: 0.9em;
}

.dropbtn {
  background-color: transparent;
  color: white;
  font-weight: bold;
  text-decoration: none;
  cursor: pointer;
}

.dropdown-content {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
  margin: 0.8em 1em;
  z-index: 1;
  right: 0;
}
.dropdown-content a {
  color: black;
  padding: 0.8em 1em;
  text-decoration: none;
  display: block;
}
.dropdown-content a:hover {background-color: #f1f1f1}
.dropdown:hover .dropdown-content {
  display: block;
}
.dropdown:hover {
  background-color: transparent;
  padding-bottom: 0.6em;
  border-bottom: 0.2em solid #f7e200;
}
.dropdown-content a.active {
	padding-bottom: 0.6em;
	border-bottom: 0.2em solid #f7e200;
}
.dropdown-content a:hover {
	padding-bottom: 0.6em;
	border-bottom: 0.2em solid #f7e200;
}
.dropbtn a.active {
	padding-bottom: 0.6em;
	border-bottom: 0.2em solid #f7e200;
}
</style>
