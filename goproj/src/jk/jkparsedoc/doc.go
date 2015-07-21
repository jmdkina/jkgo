package jkparsedoc

// Header comment define format is defined here.
// We are sorry we only support c style header file as our defined comment style.

// We support enum, #define, struct, and other function
// WARN: we only take effect if the define with a comment before. We will
// ignore it if it hasn't comment.

// Each comment must before the define.

// We support define below:
// 1. /** */ this comment is for the main define, it can be none exist.
//    It contain example create time, author, description, anything others.
// 2. /* */ This comment is for each define.
// 3. //  This define also for each define, just same with /* */.
//    If you use // you must remember take a over mark, this mark is space line or
//    //>, it will ignore the next line, if you forget it.
// 4. For now, we only support these comment
// Other comment will ignore, and also ignore with no comment define.

// We advise use /* */ for function define
// And // for other define

// All file save to bvdoc start dir
// Every dir need user create it self, we use it for user define what to parse
// from 2015-07-21 we support create dir itsself, caller need create first,
// and caller can save to the position where want.

// Some example below:
// Define enum
//   // This is enum comment
//   // This is enum comment second line
//   //> and /// will be act as the last line forced. (optional)
//   enum { }; // The real enum
//   It will find until the mark } and enum parse over.

// Define typedef
//   // This is typedef comment
//   // This is typedef comment second line
//   //> and /// will be act as the last line forced. (optional)
//   typedef struct xxx; // It will over here, if it hasn't }
//   It will find until the mark } and typedef parse over.

// Define other just like up.
//   /* This is comment start
//    * This is comment second line
//    *
//    */ This is the last comment line
//   int xxx(); // read function start.
//   It will find until the mark ; and function parse over.

// Define comment of the total header file
//   /** This is the start
//    * This is the second
//    *
//    */ The last comment line
//   It will take this to the main comment of the file.

// We only support these comment style.

// TODO:
// 1. comment and define must different.
// 2. display diff of comment and define
