// This is the 'raw' firebase setup. Outside this package, most users will want
// to use the appropriate wrapper interface, though it's OK to access these
// directly if necessary.
import firebase from 'firebase/app'
import 'firebase/app'
import 'firebase/auth'
import 'firebase/database'

let start = true;
[
  'REACT_APP_FIREBASE_API_KEY',
  'REACT_APP_FIREBASE_AUTH_DOMAIN',
  'REACT_APP_FIREBASE_DB_URL',
  'REACT_APP_FIREBASE_PROJECT_ID',
  'REACT_APP_FIREBASE_PROJECT_ID',
  'REACT_APP_FIREBASE_STORAGE_BUCKET',
  'REACT_APP_FIREBASE_MESSAGING_SENDER_ID',
].forEach(param => {
  if (process.env[param] === undefined) {
    console.error(`Did not find expected configuration parameter '${param}'. Add it to the appropirate '.env' file.`)
    start = false
  }
})

const firebaseConfig = {
  apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
  authDomain: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
  databaseURL: process.env.REACT_APP_FIREBASE_DB_URL,
  projectId: process.env.REACT_APP_FIREBASE_PROJECT_ID,
  storageBucket: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.REACT_APP_FIREBASE_MESSAGING_SENDER_ID,
}

if (!firebase.apps.length && start) {
  firebase.initializeApp(firebaseConfig)
}

export const auth = start && firebase.auth()
export const db = start && firebase.database()
