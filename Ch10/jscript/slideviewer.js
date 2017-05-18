// var touchy=('ontouchstart' in document.documentElement)?true:false;`
var RANDOM = 1;
var SEQUENTIAL = 2;
var ENGLISH_HALF = 3;
var CHINESE_HALF = 4;
var RANDOM_HALF = 5;

var SHOW_HALF = 1;
var SHOW_FULL = 2;

// overridden by form fields
var cardOrder = SEQUENTIAL;
var showHalfCard = RANDOM_HALF;

var cardState = SHOW_HALF;

var currentCard;
var cards;

// window.onkeypress = keypress;
//document.onkeypress = keypress;

function initSlides() {
    cards = new Cards(cardOrder, showHalfCard);
    cardStr = new String("card")
	
    var allDivElmts = document.getElementsByTagName("div");
    var j = 0;
    for (var i = 0; i < allDivElmts.length; i++) {
	
	if (allDivElmts[i].className == "card") {
	    cards.push(allDivElmts[i]);
	}
    }
}

function showSlides() {
    
    initSlides();
    
    allVisible = false;

    currentCard = cards.getNextCard();
    currentCard.showHalf();
}

function keypress(evt) {
    // FF var keychar = (evt.keycode) ? evt.keycode : evt.which;
    //var keychar = evt.charCode || evt.keyCode;
    keychar = evt.which;
    alert(keychar);
    if (keychar != 32) {
	return true;
    }
    
    if (cardState == SHOW_HALF) {
	currentCard.show();
	cardState = SHOW_FULL;
    } else {
	currentCard.hide();
	currentCard = cards.getNextCard();
	currentCard.showHalf();
	cardState = SHOW_HALF;
    }
    return false;
}

function advanceCard() {    
    if (cardState == SHOW_HALF) {
	currentCard.show();
	cardState = SHOW_FULL;
    } else {
	currentCard.hide();
	currentCard = cards.getNextCard();
	currentCard.showHalf();
	cardState = SHOW_HALF;
    }
}

$(document).keypress(function(evt) {
    keychar = evt.which;
    // alert(keychar);
    if (keychar != 32) {
	return true;
    }

    advanceCard();
    return false;
  }
);

/*
$(document).click(function(evt) {
    // needed for iPad, Android with no keyboard!
    advanceCard();
    return false;
  }
);
*/

document.ontouchend = function(evt) {
    // needed for iPad, Android with no keyboard!
    advanceCard();
    return false;
  }

function Cards(cardOrder, showHalfCard) {
    this.cardOrder = cardOrder;
    this.showHalfCard = showHalfCard;
    this.length = 0;
    this.index = -1;

    this.getNextCard = function() {
	if (this.cardOrder == RANDOM) {
	    index = Math.floor(Math.random() * this.length);
	    return this[index];
	}
	if (++this.index >= this.length) {
	    this.index = 0;
	}
	return this[this.index];
    }

    this.push = function(node) {
	card = new Card(node);
	this[this.length++] = card;
    }

}

function Card(node) {
    this.node = node;
    this.node.style.visibility = 'hidden';
    this.node.style.position = 'absolute';
    this.node.style.left = 0;
    this.node.style.right = 0;

    this.getChild = function(child) {
	children = this.node.getElementsByTagName("div");
	for (var i = 0; i < children.length; i++) {
	    if (children[i].className == child) {
		return children[i];
	    } 
	}
	return null;
    }

    this.english = this.getChild("english");
    this.translations = this.getChild("translations");
    this.pinyin = this.getChild("pinyin");
    this.simplified = this.getChild("simplified");
    this.traditional = this.getChild("traditional");

    this.show = function() {
	//alert("Showing full")
	 this.node.style.visibility = 'visible';
	 this.setVisibility([this.pinyin, 
			     this.simplified,
			     this.traditional],
			    'visible');
	 this.setVisibility([this.english, 
				    this.translations],
				   'visible');

    }

    this.showChild = function(child) {
	child.style.visibility = 'visible';
    }

    this.hideChild = function(child) {
	child.style.visibility = 'hidden';
    }

    this.showHalf = function() {
	if (cards.showHalfCard == RANDOM_HALF) {
	    if (Math.random() < 0.5) {
		this.setVisibility([this.pinyin, 
				    this.simplified,
				    this.traditional],
				   'hidden');
		this.setVisibility([this.english, 
				    this.translations],
				   'visible');
	    } else{ 
		this.setVisibility([this.pinyin, 
				    this.simplified,
				    this.traditional],
				   'visible');
		this.setVisibility([this.english, 
				    this.translations],
				   'hidden');
	    }
	} else if (cards.showHalfCard == CHINESE_HALF) {
		this.setVisibility([this.pinyin, 
				    this.simplified,
				    this.traditional],
				   'visible');
		this.setVisibility([this.english, 
				    this.translations],
				   'hidden');
	} else {
	    this.setVisibility([this.pinyin, 
				this.simplified,
				this.traditional],
			       'hidden');
	    this.setVisibility([this.english, 
				this.translations],
			       'visible');
	}
	this.node.style.visibility = 'visible';
    }

    this.hide = function() {
	 this.node.style.visibility = 'hidden';
	 this.setVisibility([this.pinyin, 
			     this.simplified,
			     this.traditional],
			    'hidden');
	 this.setVisibility([this.english, 
			     this.translations],
			    'hidden');
    }

    this.setVisibility = function(nodes, state) {
	for (n = 0; n < nodes.length; n++) {
	    nodes[n].style.visibility = state;
	}
    }
}
