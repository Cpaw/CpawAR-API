<template>
  <header>
    <nav>
      <a>
        <router-link tag="li" :to="{ name: 'home'}">
          Home
        </router-link>
      </a>
      <a>
        <router-link tag="li" :to="{ name: 'challenges'}">
          Challenges
        </router-link>
      </a>
      <a>
        <router-link tag="li" :to="{ name: 'ranking'}">
          Ranking
        </router-link>
      </a>
      <a>
        <router-link tag="li" :to="{ path: '/auth'}">
          Auth
        </router-link>
      </a>
    </nav>
  </header>
</template>
<script>
import axios from 'axios'

export const HTTP = axios.create({
  baseURL: `/api/v1/`
})

export default {
}
</script>
<style scoped>
nav {
  display: flex;
  justify-content: space-between;
  padding: 0 10vw;

}
a {
  text-decoration: none;
  cursor: pointer;
}
li {
  list-style-type: none;
  background-color: #fff;
  width: 14vw;
  height: 42px;
  padding-top: 6px;
  border: solid 3px #6699cc;
  border-radius: 10px 10px;
  font-size: 36px;
  font-family: "gooddog-new";
  color: #6699cc;
}
li:hover {
  border-top: solid 4px #6699cc;
  border-bottom: solid 2px #6699cc;
}
li:active {
  background-color: #fcfcfc;
}
</style>
