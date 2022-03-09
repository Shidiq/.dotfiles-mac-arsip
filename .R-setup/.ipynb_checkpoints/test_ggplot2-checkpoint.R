library(ggplot2)

pdf("out.pdf")
ggplot(diamonds) +
    geom_point(aes(x = carat, y = price, color = cut)) +
    geom_smooth(aes(x = carat, y = price, color = cut))
dev.off()
