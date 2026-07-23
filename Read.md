# Text Analyzer in Go

A simple command-line text analyzer built with Go.

## Features
- Word count
- Character count (with/without spaces)
- Vowels & Consonants
- Upper & Lower case
- Punctuations
- Numbers
- Sentences
- Lines & Spaces



Read File
   ↓
Convert to string
   ↓
Split into chunks
   ↓
Run goroutines
   ↓
Each chunk Analyze
   ↓
Send result via channel
   ↓
Merge all results
   ↓
Print final result





// Buffered (The Store-and-Forward Method)
// Unbuffered (The Direct Hand-off)

// Buffered aur Unbuffered ka sabse main simple difference yeh hai ki buffered ek "waiting area" (temporary storage) deta hai, jisse speed match ho jaati hai; jabki unbuffered mein data direct transfer hota hai aur sender-receiver ko "live" ready rehna padta

