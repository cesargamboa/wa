# 💬 WA

## The problem

Whenever you want to send a WhatsApp message to something you don't have in your contacts list,
the first step is to add him to your contacts, then open WhatsApp and wait for the contacts list
to refresh, search the user and write the message.

## This solution

`wa` its a simple binary that gets installed in your path and enables you to open the WhatsApp web
with a new chat to any number you provide and optionally with a text.

## Installation

```
curl -sf https://gobinaries.com/iamkevinwolf/wa | sh
```

## Usage

```
wa <phone number> [text]
```

### Send a WhatsApp to `000000000` number

```
wa 000000000
```

### Send a WhatsApp to `000000000` number with a predefined message

```
wa 000000000 "Hi, how are you?"
```
