<template>
  <article v-if="isSignedIn">
    <section>
      <table>
        <thead>
          <tr>
            <th>ranking</th>
            <th>name</th>
            <th>score</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in ranking" :key="user.rank">
            <td>{{user.rank}}</td>
            <td>{{user.username}}</td>
            <td>{{user.score}}</td>
          </tr>
        </tbody>
      </table>
    </section>
  </article>
  <article v-else>
    <h2>ログインしてください</h2>
  </article>
</template>
<script>
import {HTTP} from './Header'
export default {
  data () {
    return {
      ranking: {
        username: '',
        score: '',
        rank: ''
      },
      isSignedIn: false
    }
  },
  created () {
    HTTP.get(`ranking`,
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        this.$data.isSignedIn = true
        this.$data.ranking = response.data.results.ranking
      })
      .catch(e => {
        this.$data.isSignedIn = false
      })
  }
}
</script>
<style scoped>
h2 {
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 42px;
}
table {
  margin: 5vh auto;
  width: 60vw;
  border-radius: 15px;
}
thead {
  font-size: 48px;
}
tbody {
  font-size: 36px;
}
th {
  margin: 5vh auto;
  padding: 10px 0.5em 10px 0.5em;
  border-bottom: 3px solid #6699cc;
}
tr {
  margin: 5vh auto 15vh auto;
}
td {
  padding: 2vh;
}
tr th:nth-of-type(1) {
  width: 25%;
}
tr td:nth-of-type(1) {
  width: 25%;
}
tr th:nth-of-type(2) {
  width: 45%;
}
tr td:nth-of-type(2) {
  width: 45%;
}
tr th:nth-of-type(3) {
  width: 30%;
}
tr td:nth-of-type(3) {
  width: 30%;
}
</style>
