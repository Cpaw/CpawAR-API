<template>
  <article v-if="isAdmin">
    <h2>問題編集</h2>
    <section v-for="(challenge, idx) in challenges" :key="challenge.id">
      <h3>ChallengeID: {{challenge.id}}</h3>
      <form v-on:submit.prevent="editChallenge(idx)">
        <div class="field">
          <div class="label">
            <label for="title">Title:</label>
          </div>
          <div class="input">
            <input class="title" type="text" name="title" v-model="challenge.title">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="contents">Challenge text:</label>
          </div>
          <div class="contents">
            <textarea type="text" name="contents" v-model="challenge.questiontext" cols="60" rows="5" />
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="weight">Weight:</label>
          </div>
          <div class="weight">
            <input type="number" step="any" name="weight" v-model.number="challenge.Weight" />
          </div>
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
      <div class="error" v-if="isError">Invalid data</div>
      <div class="isSuccess" v-if="isSuccess">Invalid data</div>
    </section>
  </article>
</template>

<script>
import {HTTP} from '../Header'

export default {
  data () {
    return {
      challenges: [{
        id: 0,
        title: '',
        created_at: '',
        updated_at: '',
        questiontext: '',
        Weight: 0
      }],
      isSuccess: false,
      isError: false,
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
    HTTP.get(`challenges`,
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        this.$data.challenges = response.data.results.challenges
      })
  },
  methods: {
    editChallenge: function (id) {
      HTTP.put(`challenges/` + String(this.$data.challenges[id].id),
        this.$data.challenges[id],
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
          },
          withCredentials: true
        })
        .then(response => {
          this.$data.isSuccess = true
          this.$data.isError = false
        })
        .catch(e => {
          this.$data.isSuccess = false
          this.$data.isError = true
        })
    }
  }

}
</script>

<style scoped>
section {
  background: white;
  width: 60vw;
  margin: 0 auto 1em auto;
  padding: 1em;
  border: solid 3.15px #6699cc;
  border-radius: 10px 10px;
}
button {
  position: relative;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #FFF;
  background: #6699cc;
  border: solid 2px #6699cc;
  font-size: 20px;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  width: 10vw;
}
button:hover {
  cursor: pointer;
  background: #76a9dc;
}
button:active {
  background: #4679ac;
}
input, textarea {
  font-size: 16px;
}
.field {
  margin: 1em auto;
}
.label {
  font-size: 24px;
}
</style>
