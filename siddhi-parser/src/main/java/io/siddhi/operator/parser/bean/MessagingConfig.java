package io.siddhi.operator.parser.bean;

import com.google.gson.annotations.SerializedName;

public class MessagingConfig {

    @SerializedName("cluster.id")
    private String clusterId;

    @SerializedName("bootstrap.servers")
    private String[] bootstrapServers;

    public String getClusterId() {
        return clusterId;
    }

    public void setClusterId(String clusterId) {
        this.clusterId = clusterId;
    }

    public String[] getBootstrapServers() {
        return bootstrapServers;
    }

    public void setBootstrapServers(String[] bootstrapServers) {
        this.bootstrapServers = bootstrapServers;
    }

    public MessagingConfig(String clusterId, String[] bootstrapServers) {
        this.clusterId = clusterId;
        this.bootstrapServers = bootstrapServers;
    }
}