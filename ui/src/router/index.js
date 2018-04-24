import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Challenges from '@/components/Challenges'
import Challenge from '@/components/Challenge'
import Ranking from '@/components/Ranking'
import Auth from '@/components/Auth'
import Admin from '@/components/Admin'
import AddChallenge from '@/components/admin/AddChallenge'
import EditChallenge from '@/components/admin/EditChallenge'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/challenges',
      name: 'challenges',
      component: Challenges
    },
    {
      path: '/challenge/:challenge_id',
      name: 'challenge',
      component: Challenge
    },
    {
      path: '/ranking',
      name: 'ranking',
      component: Ranking
    },
    {
      path: '/auth',
      name: 'auth',
      component: Auth
    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin
    },
    {
      path: '/admin/addchallenge',
      name: 'addchallenge',
      component: AddChallenge
    },
    {
      path: '/admin/editchallenge',
      name: 'editchallenge',
      component: EditChallenge
    },
  ]
})
