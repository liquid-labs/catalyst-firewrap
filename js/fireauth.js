// 'auth' and 'db' are the underlying, raw firebase clients. Note, our exported
// methods follow (but are not always exactly the saame as) the underlying
// interface, and sometimes wrap a little extra business logic sometimes adding
// a little extra.
import { auth, db } from './firebase'

export const createUserWithEmailAndPassword = (email, password, username) =>
  auth.createUserWithEmailAndPassword(email, password)
    .then(response => {
      if (!response) {
        console.log("Response: ", response)
        throw new Error("Did not get expected response.")
      }
      const userData = response.user
      if (!userData || !userData.uid) {
        console.log("Response: ", response)
        throw new Error("Did not get expected response.")
      }
      return db.collection("users").doc(userData.uid).set({
        username: username
      })
    })

export const signInWithEmailAndPassword = (email, password) =>
  auth.signInWithEmailAndPassword(email, password)

export const signOut = () => auth.signOut()

export const sendPasswordResetEmail = (email) =>
  auth.sendPasswordResetEmail(email)

// Note, the interface for this one is a little different.
export const updatePassword = (password) =>
  auth.currentUser.updatePassword(password)

export const onAuthStateChanged = (handler/*(authUser) - null on sign out*/) =>
  auth.onAuthStateChanged(handler)
