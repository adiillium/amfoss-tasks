require 'open-uri'
require 'nokogiri'
doc = Nokogiri::HTML(open("http://www.google.com/search?q=Linux&num=10"))
doc.css('div.BNeawe.UPmit.AP7Wnd').each do |link|
    text = link.text
    text = text.gsub(' › ', '/')
    puts text
end