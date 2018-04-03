<template>
  <div id="app">
    <Header/>
    <router-view/>
  </div>
</template>
<script>
import Header from './components/Header'

export default {
  name: 'App',
  components: {
    Header
  }
}
</script>

<style>
html {
  background: url('./assets/background.png') fixed;
  background-size: cover;
}
#app {
  font-family: 'gooddog-new', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}
</style>
