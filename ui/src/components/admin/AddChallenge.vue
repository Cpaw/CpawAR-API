<template>
  <article v-if="isAdmin">
    <section class="challenge">
      <h2>問題追加</h2>
      <form v-on:submit.prevent="addChallenge">
        <div class="field">
          <div class="label">
            <label for="title">Title:</label>
          </div>
          <div class="input">
            <input class="title" type="text" name="title" v-model="challenge.Title">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="contents">Challenge text:</label>
          </div>
          <div class="contents">
            <textarea type="text" name="contents" v-model="challenge.QuestionText" cols="60" rows="5" />
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="weight">Weight:</label>
          </div>
          <div class="weight">
            <input type="number" name="weight" v-model="challenge.Weight" />
          </div>
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
      <div class="error" v-if="addChallengeError">Invalid data</div>
    </section>
  </article>
</template>

<script>
import {HTTP} from '../Header'

export default {
  data () {
    return {
      challenge: {
        Title: '',
        QuestionText: '',
        Weight: 0
      },
      addChallengeError: false,
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
        if (response.data.results == 'admin') {
          this.$data.isAdmin = true
        }
      })
      .catch(e => {
      })
  },
  methods: {
    addChallenge: function () {
      HTTP.post('challenges',
        this.$data.challenge,
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
          },
          withCredentials: true
        })
        .then(response => {
        })
        .catch(e => {
          this.$data.addChallengeError = true
        })
    }
  }
}
</script>

<style scoped>
article {
  display: -webkit-flex;
  display: flex;
}
section {
  -webkit-justify-content: space-around;
  justify-content: space-around;
  margin: 0 auto;
}
label {
  width: 8vw;
}
input, button {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 20px;
  border: solid 2px #6688cc;
  border-radius: 5px;
  outline: 0;
}
input:focus {
  border: solid 2px #6699cc;
  border-radius: 5px;
}
textarea {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 20px;
  border: solid 2px #6688cc;
  border-radius: 5px;
  outline: 0;
}
button {
  position: relative;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #fff;
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
.field {
  margin: 0 auto 5vh auto;
  padding: 5px;
  font-size: 24px;
  font-weight: 700;
  width: 80vw;
}
.contents {

}
.label {
  text-align: center;
}
.error {
  margin: 2vh auto 2vh auto;
  font-size: 24px;
  color: #ff5d86;
}
</style>
