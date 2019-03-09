// 'auth' is the underlying firebase client. Note, our exported methods follow
// but are not always exactly the saame as the underlying interface.
import { auth } from './firebase'

const createUserWithEmailAndPassword = async(email, password, displayName) => {
  const userCredentials = await
    auth.createUserWithEmailAndPassword(email, password)
  if (displayName) {
    const authUser = userCredentials.user
    await authUser.updateProfile({displayName : displayName})
  }
  console.log('userCredentials now: ', userCredentials)
  return userCredentials
}

const signInWithEmailAndPassword = (email, password) =>
  auth.signInWithEmailAndPassword(email, password)

// Note we use Catalyst standard 'log out'
const logOut = () => auth.signOut()

const sendPasswordResetEmail = (email) =>
  auth.sendPasswordResetEmail(email)

// Note, the interface for this one is a little different.
const updatePassword = (password) =>
  auth.currentUser.updatePassword(password)

const onAuthStateChanged = (handler/*(authUser) - null on sign out*/) =>
  auth.onAuthStateChanged(handler)


export {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  logOut,
  sendPasswordResetEmail,
  updatePassword,
  onAuthStateChanged,
}
