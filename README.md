kraken
========

Automated anti-phishing tool for dastardly form-based mayhem

## Description

Damn phishers! They fooled Gammy again with some kind of "change your bank password, your account is at risk!" bullshit, and now there's a significant chunk of your inheritance missing!  Oh, that makes you so mad.  You're so mad you could spit, but what can you do? Those stupid little form-based scams they put up are a dime-a-dozen! Sure, someone eventually takes them down, but by that point they've got plenty of actionable data, and they can just put up another one! Man it sure is awful that no one can do anything about this. What has the world come to?

Kraken is a simple tool designed to parse a form url[s] and submit millions and millions of fake form submissions to the attacker, hopefully drowning them in so much fake data that they can't make use of any of it any more. Kraken will parse forms and attempt to do its damnedest to pick the right fake-ass entries to put in the right fake-ass form fields before you tell it to just go to town. It will then gleefully report how many thousands and millions of fake entries it has made until you tell it to stop or (ideally) the scammer finds out what is going on, wails unto the heavens at the cruel impassiveness of the intertubes, and packs up his-or-her bag of snake oil and sails off into the sunset where they are hopefully eaten by sharks.

## Installation

These docs are shitty.  They are also new, and will be less shitty very soon.

```bash
$ git clone https://github.com/anaxagoras/kraken.git
$ cd wherever/you/just/put/that
$ go build
```

## Running

```bash
$ ./kraken
```

## Todo

- Pretty much everything form-related.
- Address generation with valid zip codes
- Non-US valid address generation
