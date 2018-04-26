<template>
  <article v-if="isAdmin">
    <section class="challenge">
      <router-link tag="li" :to="{ name: 'addchallenge' }">Add Challenge</router-link>
      <router-link tag="li" :to="{ name: 'editchallenge' }">Edit Challenge</router-link>
    </section>
     <section class="user">
      <router-link :to="{ name: 'edituser' }">Edit User</router-link>
    </section>
  </article>
</template>

<script>
import {HTTP} from './Header'

export default {
  data () {
    return {
      isAdmin: false
    }
  },
  mounted () {
    HTTP.get(`role`,
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        if (response.data.results === 'admin') {
          this.$data.isAdmin = true
        }
      })
      .catch(e => {
      })
  }
}
</script>

<style scoped>
article {
  width: 20vw;
  margin: 5vw auto;
  display: flex;
  justify-content: space-between;
}
li {
  list-style-type: none;
  background-color: #fff;
  width: 14vw;
  height: 42px;
  padding-top: 6px;
  margin: 1em;
  border: solid 3px #6699cc;
  border-radius: 10px 10px;
  font-size: 36px;
  font-family: "gooddog-new";
  color: #6699cc;
}
li:hover {
  border-top: solid 4px #6699cc;
  border-bottom: solid 2px #6699cc;
  cursor: pointer;
}
li:active {
  background-color: #fcfcfc;
}
</style>
