/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-06
 * @fileoverview This program asks the user for their first name and then
 * displays the lyrics to the song "Happy Birthday," inserting the user's
 * name at the appropriate place in the song.
 */

// variables
let userName: string;

// input
userName = prompt("What is your first name?") || "Friend";

// output
console.log("\n");
console.log("Here is your birthday song:\n");
console.log("Happy Birthday to you,");
console.log("Happy Birthday to you,");
console.log("Happy Birthday dear " + userName + ",");
console.log("Happy Birthday to you!");
console.log("\nDone.");

