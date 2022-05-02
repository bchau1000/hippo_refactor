import { initializeApp, type FirebaseApp } from "firebase/app";
import { getAuth, type Auth } from "firebase/auth";

const firebaseConfig:Object = {
  apiKey: "AIzaSyC3mclhmR80s6gReKsWXPVoUA7sB_EOMSw",
  authDomain: "hippo-3b249.firebaseapp.com",
  projectId: "hippo-3b249",
  storageBucket: "hippo-3b249.appspot.com",
  messagingSenderId: "264264475909",
  appId: "1:264264475909:web:db138310da6a959c5ec36c",
  measurementId: "G-D99EENT984"
};

export const app:FirebaseApp = initializeApp(firebaseConfig);
export const auth:Auth = getAuth();